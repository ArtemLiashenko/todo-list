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
	Route{ // Complete
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{ // Complete
		"TodoIndex",
		"GET",
		"/todos",
		TodoIndex,
	},
	Route{ // Complete
		"TodoShow",
		"GET",
		"/todos/{todoId}",
		TodoShow,
	},
	Route{
		"TodoCreate",
		"POST",
		"/todos",
		TodoCreate,
	},
	Route{
		"TodoShow",
		"DELETE",
		"/todos/{todoId}",
		TodoDelete,
	},
}
