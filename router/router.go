package router

import (
	"net/http"
	"projet_wiki/controllers"

	"github.com/gorilla/mux"
)

//This is a router, powered by mux!
//The route type contains a name for the route, a method (PUT, GET, POST...) a pattern (the url, basically) and a handler function that
//makes the connection between a controller and the route.

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

// NewRouter is the register for every public route
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	//Check every route created below register it
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

//This is gonna be parsed, so this must contain every public route.
var routes = Routes{
	Route{
		Name:        "hello",
		Method:      "GET",
		Pattern:     "/hello",
		HandlerFunc: controllers.HelloHandler,
	},
	Route{
		Name:        "displayData",
		Method:      "GET",
		Pattern:     "/displayData",
		HandlerFunc: controllers.DisplayData,
	},
	Route{
		Name:        "getData",
		Method:      "GET",
		Pattern:     "/data",
		HandlerFunc: controllers.GetData,
	},
	Route{
		Name:        "post example",
		Method:      "POST",
		Pattern:     "/data/received",
		HandlerFunc: controllers.ReceiveData,
	},
}
