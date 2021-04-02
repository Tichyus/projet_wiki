package router

import (
	"net/http"
	"projet_wiki/controllers"
	"projet_wiki/middleware"

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

	// Authentication routes

	Route{
		Name:        "Sign in",
		Method:      "POST",
		Pattern:     "/signin",
		HandlerFunc: controllers.Signin,
	},
	Route{
		Name:        "Sign up",
		Method:      "POST",
		Pattern:     "/signup",
		HandlerFunc: controllers.CreateUser,
	},
	Route{
		Name:        "Refresh jwt",
		Method:      "POST",
		Pattern:     "/refreshToken",
		HandlerFunc: controllers.RefreshToken,
	},

	// User management routes

	Route{
		Name:        "Create user",
		Method:      "POST",
		Pattern:     "/user/create",
		HandlerFunc: middleware.VerifyJwt(controllers.CreateUser),
	},

	Route{
		Name:        "Read user",
		Method:      "GET",
		Pattern:     `/user/{ID}`,
		HandlerFunc: controllers.ReadUser,
	},

	Route{
		Name:        "Update user",
		Method:      "POST",
		Pattern:     "/user/update",
		HandlerFunc: middleware.VerifyJwt(controllers.UpdateUser),
	},

	Route{
		Name:        "Delete user",
		Method:      "POST",
		Pattern:     `/delete/delete`,
		HandlerFunc: middleware.VerifyJwt(controllers.DeleteUser),
	},

	// Articles management routes

	Route{
		Name:        "list all articles",
		Method:      "GET",
		Pattern:     "/article",
		HandlerFunc: controllers.AllArticles,
	},

	Route{
		Name:        "list all articles from one user",
		Method:      "GET",
		Pattern:     `/user/{ID}/articles`,
		HandlerFunc: controllers.AllArticlesFromUser,
	},

	Route{
		Name:        "Read specific article ",
		Method:      "GET",
		Pattern:     "/article/{ID}",
		HandlerFunc: controllers.ReadArticle,
	},

	Route{
		Name:        "create article",
		Method:      "POST",
		Pattern:     `/article/create`,
		HandlerFunc: middleware.VerifyJwt(controllers.CreateArticle),
	},

	Route{
		Name:        "delete article",
		Method:      "POST",
		Pattern:     "/article/delete",
		HandlerFunc: middleware.VerifyJwt(controllers.DeleteArticle),
	},

	Route{
		Name:        "Update article",
		Method:      "POST",
		Pattern:     `/article/update`,
		HandlerFunc: middleware.VerifyJwt(controllers.UpdateArticle),
	},

	// Comment management routes

	Route{
		Name:        "Read specific comment",
		Method:      "GET",
		Pattern:     "/comment/{ID}",
		HandlerFunc: controllers.ReadComment,
	},

	Route{
		Name:        "list all comments from one article",
		Method:      "GET",
		Pattern:     `/articles/{ID}/comments`,
		HandlerFunc: controllers.ReadComments,
	},

	Route{
		Name:        "create comment",
		Method:      "POST",
		Pattern:     "/comment/create",
		HandlerFunc: middleware.VerifyJwt(controllers.CreateComment),
	},

	Route{
		Name:        "delete comment",
		Method:      "POST",
		Pattern:     `/comment/delete`,
		HandlerFunc: middleware.VerifyJwt(controllers.DeleteComment),
	},

	Route{
		Name:        "update comment",
		Method:      "POST",
		Pattern:     "/comment/update",
		HandlerFunc: middleware.VerifyJwt(controllers.UpdateComment),
	},

	// Simple 

	Route{
		Name:        "Welcome user",
		Method:      "GET",
		Pattern:     "/welcome",
		HandlerFunc: middleware.VerifyJwt(controllers.Welcome),
	},
}
