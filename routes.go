package main

import "net/http"

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
		"SensuCheckResults",
		"GET",
		"/sensu/checks/{clientName}",
		SensuChecksToHtml,
	},
	Route{
		"AlwaysOK",
		"GET",
		"/alwaysok",
		RouteOK,
	},
}
