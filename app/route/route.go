package route
 
import (
	"github.com/gorilla/mux"
	"projet_wiki/app/controllers"
)

func GetRoutes() *mux.Router {
    routes := mux.NewRouter().StrictSlash(false)
 
    routes := routes.PathPrefix("/api/v1/").Subrouter()
 
    routes.HandleFunc("/hello", controllers.Hello).Methods("GET")
 
    return routes;
}