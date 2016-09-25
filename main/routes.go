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
		index,
	},
	Route{
		"TodoIndex",
		"GET",
		"/todos",
		todoIndex,
	},
	Route{
		"TodoShow",
		"GET",
		"/todos/{todoId}",
		todoShow,
	},
	Route{
		"TodoCreate",
		"POST",
		"/todos",
		todoCreate,
	},
}
