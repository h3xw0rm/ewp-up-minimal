package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type routes []route

func newRouter() *mux.Router {

	CatalogueGlob = getCatalog()

	router := mux.NewRouter().StrictSlash(true)

	rest := router.PathPrefix("/rest").Subrouter()
	for _, route := range restRoutes {
		rest.
			HandleFunc(route.Pattern, route.HandlerFunc).
			Name(route.Name)
	}

	return router
}

var restRoutes = routes{
	route{
		"manifest",
		"GET",
		"/manifest",
		createManifest,
	},
	route{
		"echo",
		"",
		"/echo",
		echoRestHandler,
	},
	route{
		"institutions",
		"",
		"/institutions",
		institutionsRestHandler,
	},
}
