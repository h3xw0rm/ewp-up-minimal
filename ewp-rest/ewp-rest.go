package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	//MaxHeiIDs allowed to query
	MaxHeiIDs = 1
)

var nonce = map[string]time.Time{}

//CatalogueGlob Reader from Registry
var CatalogueGlob = &catalogue{}

//AcceptedInstitutions is a list of Institutions Accepted
var AcceptedInstitutions = []string{os.Getenv("EWP_SCHAC")}

/*Hostname DNS*/
var Hostname = os.Getenv("EWP_DNS")

/*Port for connections*/
var Port = os.Getenv("EWP_REST_INT_PORT")

/*PrivKey for server*/
var PrivKey = "/go/certs/" + os.Getenv("EWP_CERT_KEY")

/*PubKey for server*/
var PubKey = "/go/certs/" + os.Getenv("EWP_CERT_PUB")

/*RawPubKey for server*/
var RawPubKey = "/go/certs/" + os.Getenv("EWP_CERT_PUBRAW")

/*RSAPubKey for server*/
var RSAPubKey = "/go/certs/" + os.Getenv("EWP_CERT_RSAPUB")

/*RawRSAPubKey for server*/
var RawRSAPubKey = "/go/certs/" + os.Getenv("EWP_CERT_RSAPUBRAW")

/*PubChain of DigiCert for server*/
var PubChain = "/go/certs/" + os.Getenv("EWP_CERT_CHAIN")

/*PubAll of server plus DigiCert*/
var PubAll = "/go/certs/" + os.Getenv("EWP_CERT_CHAINALL")

func main() {

	logPath := "/data/" + os.Getenv("EWP_LOGFILE_SERVER")

	lf, err := os.OpenFile(logPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0640)

	if err != nil {
		log.Fatal("OpenLogfile: os.OpenFile:", err)
	}

	log.SetOutput(lf)

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// call function to create the routes
	r := newRouter()

	clientCertPool := x509.NewCertPool()

	clientCACert, err := ioutil.ReadFile(PubChain)
	if err != nil {
		log.Fatal("Unable to open cert", err)
	}
	clientCertPool.AppendCertsFromPEM(clientCACert)

	tlsConfig := &tls.Config{
		ClientAuth:         tls.RequestClientCert,
		InsecureSkipVerify: true,
	}

	tlsConfig.BuildNameToCertificate()

	server := &http.Server{
		TLSConfig:    tlsConfig,
		Addr:         "0.0.0.0:" + Port,
		Handler:      r,
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  60 * time.Second,
	}

	//start listening on port 443
	log.Fatal(server.ListenAndServeTLS(PubAll, PrivKey))
}
