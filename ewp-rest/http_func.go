package main

import (
	"bytes"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/lestrrat/go-libxml2/parser"
	xmlsec "github.com/lestrrat/go-xmlsec"
	"github.com/spacemonkeygo/httpsig"
)

func isValidUUID(uuid string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid)
}
func isUsedUUID(uuid string) bool {
	for uid, datetime := range nonce {
		if time.Now().Sub(datetime).Minutes() > 5.0 {
			delete(nonce, uid)
		}
	}

	if _, ok := nonce[uuid]; ok {
		return false
	}

	nonce[uuid] = time.Now()

	return true
}

func verifyRequestSignature(r *http.Request, w http.ResponseWriter, headers map[string]string) bool {

	heads := []string{}

	for k := range headers {
		heads = append(heads, strings.ToLower(k))
	}

	getter := setKeys()

	verifier := httpsig.NewVerifier(getter.KeyGetter)
	verifier.SetRequiredHeaders(heads)
	err := verifier.Verify(r)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

type keyStorage struct {
	httpsig.KeyGetter
}

func setKeys() *keyStorage {

	keystore := httpsig.NewMemoryKeyStore()

	for _, bin := range CatalogueGlob.Binaries {
		for _, rsaCert := range bin.RSACert {
			pubPEM := "-----BEGIN RSA PUBLIC KEY-----\n" + strings.TrimSpace(rsaCert.Value) + "\n-----END RSA PUBLIC KEY-----"
			block, _ := pem.Decode([]byte(pubPEM))
			if block == nil {
				fmt.Println("failed to parse PEM block containing the public key")
				return nil
			}

			pub, err := x509.ParsePKIXPublicKey(block.Bytes)
			if err != nil {
				fmt.Println("failed to parse DER encoded public key: " + err.Error())
				return nil
			}
			keystore.SetKey(string(rsaCert.SHA256), pub)
		}
	}

	return &keyStorage{
		KeyGetter: keystore,
	}
}

func verifyClientAuthHTTPSig(r *http.Request, w http.ResponseWriter) (bool, string, int, []catalogueHosts) {

	var hostDiscovered []catalogueHosts

	auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

	if auth[0] != "Signature" {
		return false, "Signature not present", 403, hostDiscovered
	}

	signList := strings.Split(auth[1], ",")

	sigRSA := false

	validKey := ""
	head := []string{}
	signature := ""

	var signedHeaders []string

	for _, val := range signList {
		l := strings.SplitN(val, "=", 2)
		k, v := l[0], l[1]
		v = v[1 : len(v)-1]

		if strings.Compare(strings.ToLower(k), "algorithm") == 0 && strings.Compare(strings.ToLower(v), "rsa-sha256") == 0 {
			sigRSA = true
		}

		if strings.Compare(strings.ToLower(k), "headers") == 0 {
			head = strings.Split(v, " ")
			rt := false
			ho := false
			da := false
			di := false
			xr := false
			for _, h := range head {
				signedHeaders = append(signedHeaders, h)
				if strings.Compare(strings.ToLower(h), "(request-target)") == 0 {
					rt = true
				}
				if strings.Compare(strings.ToLower(h), "host") == 0 {
					ho = true
				}
				if strings.Compare(strings.ToLower(h), "date") == 0 {
					da = true
				}
				if strings.Compare(strings.ToLower(h), "original-date") == 0 {
					da = true
				}
				if strings.Compare(strings.ToLower(h), "digest") == 0 {
					di = true
				}
				if strings.Compare(strings.ToLower(h), "x-request-id") == 0 {
					xr = true
				}
			}
			if !rt || !ho || !da || !di || !xr {
				return false, "Required headers not present", 401, hostDiscovered
			}
		}

		if strings.Compare(strings.ToLower(k), "keyid") == 0 {
			for _, host := range CatalogueGlob.Hosts {
				for _, rsaPubKey := range host.CliCredInUse.RSACert {
					if strings.Compare(v, string(rsaPubKey.SHA256)) == 0 {
						validKey = v
					}
				}
			}
		}

		if strings.Compare(strings.ToLower(k), "signature") == 0 {
			signature = v
		}
	}

	signedHeadersValues := map[string]string{}
	for _, v := range signedHeaders {
		if strings.Compare(v, "(request-target)") != 0 {
			if strings.Compare(strings.ToLower(v), "host") == 0 {
				signedHeadersValues[v] = strings.SplitN(r.Host, ":", 2)[0]
				r.Host = strings.SplitN(r.Host, ":", 2)[0]
			} else {
				if strings.Compare("", r.Header.Get(v)) == 0 {
					return false, "Some signed headers has no values: " + v, 401, hostDiscovered
				}
				signedHeadersValues[v] = r.Header.Get(v)
			}
		} // else {
		//signedHeadersValues[v] = r.Method + " " + r.URL.String()
		//}
	}

	validSig := verifyRequestSignature(r, w, signedHeadersValues)

	if !validSig {
		return false, "Signature not valid ", 401, hostDiscovered
	}

	if !sigRSA {
		return false, "Wrong signature algorithm", 403, hostDiscovered
	}

	if strings.Compare(validKey, "") != 0 && strings.Compare(signature, "") != 0 {

		if strings.Compare(strings.ToLower(r.Header.Get("Host")), Hostname) != 0 && strings.Compare(r.Header.Get("Host"), "") != 0 {
			return false, "Host not correct", 403, hostDiscovered
		}

		const longForm = "Mon, 2 Jan 2006 15:04:05 MST"

		if strings.Compare(r.Header.Get("Date"), "") != 0 {
			data, err := time.Parse(longForm, r.Header.Get("Date"))
			if err != nil {
				fmt.Println(err)
				return false, "An error occured processing data", 500, hostDiscovered
			}
			if !data.After(time.Now().Local().Add(time.Minute * (-time.Duration(5)))) {
				return false, "Date before accepted 5 minutes", 400, hostDiscovered
			}
		}
		if strings.Compare(r.Header.Get("Original-Date"), "") != 0 {
			data, err := time.Parse(longForm, r.Header.Get("Original-Date"))
			if err != nil {
				fmt.Println(err)
				return false, "An error occured processing data", 500, hostDiscovered
			}
			if !data.After(time.Now().Local().Add(time.Minute * (-time.Duration(5)))) {
				return false, "Date before accepted 5 minutes", 400, hostDiscovered
			}
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return false, "Error reading message body", 403, hostDiscovered
		}
		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		sha256value := sha256.Sum256(body)
		slicedSHA := sha256value[:]
		slicedSHAbase64 := base64.StdEncoding.EncodeToString(slicedSHA)
		digestList := strings.Split(r.Header.Get("Digest"), ",")
		validDig := false
		for _, dig := range digestList {
			digest := strings.SplitN(dig, "=", 2)
			if strings.Compare(strings.ToUpper(digest[0]), "SHA-256") == 0 && strings.Compare(digest[1], slicedSHAbase64) == 0 {
				validDig = true
			}
		}

		if !validDig {
			return false, "Wrong digest", 400, hostDiscovered
		}

		if !isValidUUID(r.Header.Get("X-Request-Id")) {
			return false, "Not valid format for X-Request-Id", 400, hostDiscovered
		}

		if !isUsedUUID(r.Header.Get("X-Request-Id")) {
			return false, "X-Request-Id already used", 400, hostDiscovered
		}

		for _, host := range CatalogueGlob.Hosts {
			for _, creds := range host.CliCredInUse.RSACert {
				if bytes.Compare([]byte(validKey), []byte(strings.ToLower(string(creds.SHA256)))) == 0 {
					hostDiscovered = append(hostDiscovered, host)
				}
			}
		}

		//if len(hostDiscovered) > 0 {
		return true, "", 0, hostDiscovered
		//}
	}

	return false, "Key or signature not present.", 403, hostDiscovered
}

func verifyClientAuthTLS(r *http.Request, w http.ResponseWriter) (bool, []catalogueHosts) {
	var hostDiscovered = []catalogueHosts{}

	for _, certs := range r.TLS.PeerCertificates {
		hash := sha256.Sum256(certs.Raw)
		h := hex.EncodeToString(hash[0:])
		for _, host := range CatalogueGlob.Hosts {
			for _, creds := range host.CliCredInUse.Certificate {
				if bytes.Compare([]byte(h), []byte(strings.ToLower(string(creds.SHA256)))) == 0 {
					hostDiscovered = append(hostDiscovered, host)
				}
			}
		}
	}

	if len(hostDiscovered) > 0 {
		return true, hostDiscovered
	}

	return false, hostDiscovered
}

func testSecurityAndGetHosts(r *http.Request, w http.ResponseWriter) (bool, string, int, []catalogueHosts) {

	log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

	var validSecurity, reason, code, hostDiscovered = verifyClientAuthHTTPSig(r, w)

	if validSecurity {
		return true, "", 0, hostDiscovered
	}

	validSecurity, hostDiscovered = verifyClientAuthTLS(r, w)
	if validSecurity {
		return true, "", 0, hostDiscovered
	}

	return false, reason, code, hostDiscovered

}

func writeResponse(r *http.Request, w http.ResponseWriter, status int, x []byte) {

	type gzipResponseWriter struct {
		io.Writer
		http.ResponseWriter
	}

	headers := map[string]string{}

	if status == http.StatusUnauthorized {
		w.Header().Set("WWW-Authenticate", "Signature realm=\"EWP\"")
		headers["WWW-Authenticate"] = "Signature realm=\"EWP\""
		w.Header().Set("Want-Digest", "SHA-256")
		headers["Want-Digest"] = "SHA-256"
	}

	var needToSign = false

	if strings.Compare(r.Header.Get("Accept-Signature"), "") != 0 {
		sigList := strings.Split(r.Header.Get("Accept-Signature"), ",")
		for _, val := range sigList {
			if strings.Compare(val, "rsa-sha256") == 0 {
				needToSign = true
			}
		}
	}

	w.Header().Set("Content-Type", "application/xml")

	xmlsec.Init()
	defer xmlsec.Shutdown()

	p := parser.New(parser.XMLParseDTDLoad | parser.XMLParseDTDAttr | parser.XMLParseNoEnt)
	doc, err := p.ParseString(string(x))
	if err != nil {
		fmt.Printf("DocumentElement failed:")
		fmt.Println(err)
	}

	if needToSign {

	}

	w.Header().Set("Host", Hostname)
	w.Header().Set("Date", time.Now().UTC().Format("Mon, 2 Jan 2006 15:04:05 GMT"))
	headers["Date"] = time.Now().UTC().Format("Mon, 2 Jan 2006 15:04:05 GMT")

	if strings.Compare(r.Header.Get("X-Request-Id"), "") != 0 {
		headers["X-Request-Id"] = r.Header.Get("X-Request-Id")
		w.Header().Set("X-Request-Id", r.Header.Get("X-Request-Id"))
	}

	w.Header().Set("Content-Length", strconv.Itoa(len([]byte(doc.Dump(true)))))

	if strings.Compare(r.Header.Get("Authorization"), "") != 0 {
		signature := strings.SplitN(r.Header.Get("Authorization"), " ", 2)[1]
		sigList := strings.Split(signature, ",")
		for _, v := range sigList {
			if strings.HasPrefix(v, "signature=") {
				headers["X-Request-Signature"] = v[11 : len(v)-1]
				w.Header().Set("X-Request-Signature", v[11:len(v)-1])
			}
		}
	}

	/*if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
		//headers["Content-Encoding"] = "gzip"
		w.Header().Set("Content-Encoding", "gzip")
		var b bytes.Buffer
		gz := gzip.NewWriter(&b)
		//w.WriteHeader(status)
		if _, err := gz.Write([]byte(doc.Dump(true))); err != nil {
			panic(err)
		}
		if err := gz.Flush(); err != nil {
			panic(err)
		}
		if err := gz.Close(); err != nil {
			panic(err)
		}
		sha256value := sha256.Sum256(b.Bytes())
		slicedSHA := sha256value[:]
		slicedSHAbase64 := base64.StdEncoding.EncodeToString(slicedSHA)
		headers["Digest"] = "SHA-256=" + slicedSHAbase64
		w.Header().Set("Digest", "SHA-256="+slicedSHAbase64)

		//dump := base64.StdEncoding.EncodeToString(b.Bytes())

		if needToSign {
			sig := signRequest(r, w, headers)
			w.Header().Set("Signature", sig)
		}
		w.WriteHeader(status)
		fmt.Println(b.Bytes())
		bodygz := gzip.NewWriter(w)
		defer bodygz.Close()
		bodygz.Write([]byte(doc.Dump(true)))
		//dec, _ := base64.StdEncoding.DecodeString(dump)
		//w.Write(dec)
	} else {*/
	sha256value := sha256.Sum256([]byte(doc.Dump(true)))
	slicedSHA := sha256value[:]
	slicedSHAbase64 := base64.StdEncoding.EncodeToString(slicedSHA)
	headers["Digest"] = "SHA-256=" + slicedSHAbase64
	w.Header().Set("Digest", "SHA-256="+slicedSHAbase64)

	if needToSign {
		sig := signRequest(r, w, headers)
		w.Header().Set("Signature", sig)
	}
	w.WriteHeader(status)

	w.Write([]byte(doc.Dump(true)))
	//}
}

func signRequest(r *http.Request, w http.ResponseWriter, headers map[string]string) string {
	newReq, err := http.NewRequest(r.Method, r.URL.String(), nil)

	heads := []string{}

	for k, v := range headers {
		newReq.Header.Set(k, v)
		heads = append(heads, strings.ToLower(k))
	}

	privContent, err := ioutil.ReadFile(PrivKey)
	if err != nil {
		log.Fatal(err)
	}
	privBlock, _ := pem.Decode(privContent)
	key, err := x509.ParsePKCS1PrivateKey(privBlock.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	rsaPubKey, err := ioutil.ReadFile(RawRSAPubKey)
	if err != nil {
		fmt.Println(err)
	}
	b64dec, err := base64.StdEncoding.DecodeString(strings.TrimSpace(string(rsaPubKey)))
	if err != nil {
		fmt.Println(err)
	}
	sha256sum := sha256.Sum256(b64dec)
	shahex := hex.EncodeToString(sha256sum[:])

	signer := httpsig.NewSigner(shahex, key, httpsig.RSASHA256, heads)
	err = signer.Sign(newReq)
	if err != nil {
		log.Printf("Signature failed: %s", err)
	}
	return strings.SplitN(newReq.Header.Get("Authorization"), " ", 2)[1]
}
