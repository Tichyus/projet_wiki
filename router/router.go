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


	// User management routes

	Route{
		Name:        "Create user",
		Method:      "POST",
		Pattern:     "/create-user",
		HandlerFunc: controllers.CreateUser,
	},

	Route{
		Name:        "Read user",
		Method:      "GET",
		Pattern:     `/read-user/{id}`,
		HandlerFunc: controllers.ReadUser,
	},

	Route{
		Name:        "Update user",
		Method:      "POST",
		Pattern:     "/update-user",
		HandlerFunc: controllers.UpdateUser,
	},

	Route{
		Name:        "Delete user",
		Method:      "POST",
		Pattern:     `/delete-user`,
		HandlerFunc: controllers.DeleteUser,
	},

	// Articles management routes

	Route{
		Name:        "list all articles",
		Method:      "GET",
		Pattern:     "/articles",
		HandlerFunc: controllers.AllArticles,
	},

	Route{
		Name:        "list all articles from one user",
		Method:      "GET",
		Pattern:     `/articles/user/{id}`,
		HandlerFunc: controllers.AllArticlesFromUser,
	},

	Route{
		Name:        "Read specific article ",
		Method:      "GET",
		Pattern:     "/articles/{id}",
		HandlerFunc: controllers.ReadArticle,
	},

	Route{
		Name:        "create article",
		Method:      "POST",
		Pattern:     `/create-article`,
		HandlerFunc: controllers.CreateArticle,
	},

	Route{
		Name:        "delete article",
		Method:      "POST",
		Pattern:     "/delete-article",
		HandlerFunc: controllers.DeleteArticle,
	},

	Route{
		Name:        "Update article",
		Method:      "POST",
		Pattern:     `/update-article`,
		HandlerFunc: controllers.UpdateArticle,
	},

	// Comment management routes

	Route{
		Name:        "Read specific comment",
		Method:      "GET",
		Pattern:     "/comments/{id}",
		HandlerFunc: controllers.ReadComment,
	},

	Route{
		Name:        "list all comments from one article",
		Method:      "GET",
		Pattern:     `/articles/{id}/comment`,
		HandlerFunc: controllers.ReadComments,
	},

	Route{
		Name:        "create comment",
		Method:      "POST",
		Pattern:     "/create-comment",
		HandlerFunc: controllers.CreateComment,
	},

	Route{
		Name:        "delete comment",
		Method:      "POST",
		Pattern:     `/delete-comment`,
		HandlerFunc: controllers.DeleteComment,
	},

	Route{
		Name:        "update comment",
		Method:      "POST",
		Pattern:     "/update-comment",
		HandlerFunc: controllers.UpdateComment,
	},
}
