package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"AdminIndex",
		"GET",
		"/admin",
		AdminIndex,
	},
	Route{
		"PurlIndex",
		"GET",
		"/purls",
		PurlIndex,
	},
	Route{
		"PurlCreate",
		"POST",
		"/purl/create",
		PurlCreate,
	},
	Route{
		"PurlShow",
		"GET",
		"/view/{purlId}",
		PurlShow,
	},
	Route{
		"PurlShowFile",
		"GET",
		"/view/{purlId}/{filename}",
		PurlShowFile,
	},
	Route{
		"Query",
		"GET",
		"/query?={query}",
		Query,
	},
}

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}

	return router
}
