package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func getCatalog() (cat *catalogue) {
	url := "https://dev-registry.erasmuswithoutpaper.eu/catalogue-v1.xml"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	err = xml.Unmarshal(body, &cat)
	if err != nil {
		fmt.Println(err)
	}

	return cat
}

func getCodeForHost(name string) string {

	for _, insts := range CatalogueGlob.Institutions.Heis {
		if strings.Compare(insts.ID, name) == 0 {
			for _, otherID := range insts.OtherID {
				if strings.Compare(otherID.Type, "erasmus") == 0 {
					return otherID.Value
				}
			}
		}
	}

	return ""
}

func checkFormValues(w http.ResponseWriter, r *http.Request) map[string][]string {
	var m map[string][]string

	err := r.ParseForm()
	if err != nil {
		fmt.Println("Error parsing")
		fmt.Println(err)
		return m
	}

	if strings.Compare("POST", strings.ToUpper(r.Method)) == 0 {
		m = r.PostForm
	} else {
		m = r.Form
	}

	return m
}

func checkFormType(w http.ResponseWriter, r *http.Request) bool {

	switch r.Method {
	case "GET":
		break
	case "POST":
		types := strings.Split(r.Header.Get("content-type"), ";")
		if strings.Compare(types[0], "application/x-www-form-urlencoded") != 0 {
			return false
		}
		break
	default:
		return false
	}
	return true
}

func getTLSConfigurations() *tls.Config {
	cert, err := tls.LoadX509KeyPair(PubKey, PrivKey)
	if err != nil {
		log.Fatalln("Unable to load cert", err)
	}
	clientCACert, err := ioutil.ReadFile(PubChain)
	if err != nil {
		log.Fatal("Unable to open cert", err)
	}
	clientCertPool := x509.NewCertPool()
	clientCertPool.AppendCertsFromPEM(clientCACert)

	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
	}

	tlsConfig.BuildNameToCertificate()

	return tlsConfig
}
