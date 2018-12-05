package main

import (
	"encoding/xml"
	"log"
	"net/http"
)

func getError40X(w http.ResponseWriter, r *http.Request, reason string, code int) {
	switch code {
	case 400:
		getError400(w, r, reason)
		break
	case 401:
		getError401(w, r, reason)
		break
	case 403:
		getError403(w, r, reason)
		break
	default:
		getError401(w, r, reason)
	}
}

func createManifest(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

	manifest := getEWPManifest()

	status := http.StatusOK

	x, err := xml.MarshalIndent(manifest, "", "  ")
	if err != nil {
		status = http.StatusInternalServerError
		getError500(w, r)
	} else {
		writeResponse(r, w, status, x)
	}
}

func echoRestHandler(w http.ResponseWriter, r *http.Request) {
	var valid, reason, code, hostsCovered = testSecurityAndGetHosts(r, w)
	if valid {
		if checkFormType(w, r) {
			queryValues := checkFormValues(w, r)
			retrieveEcho(w, r, hostsCovered, queryValues)
		} else {
			getError405(w, r)
		}
	} else {
		getError40X(w, r, reason, code)
	}
}

func institutionsRestHandler(w http.ResponseWriter, r *http.Request) {
	var valid, reason, code, hostsCovered = testSecurityAndGetHosts(r, w)
	if valid {
		if checkFormType(w, r) {
			queryValues := checkFormValues(w, r)
			retrieveInstitutions(w, r, hostsCovered, queryValues)
		} else {
			getError405(w, r)
		}
	} else {
		getError40X(w, r, reason, code)
	}
}
