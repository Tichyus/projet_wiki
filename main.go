package main


import (
	"projet_wiki/router"
	"log"
	"net/http"
)

const dwldPath = "./tmp"

func main() {

	port := "8080"
	newRouter := router.NewRouter()

	log.Print("\nServer started on port " + port)

	newRouter.PathPrefix("/files/").Handler(http.StripPrefix("/files/",
		http.FileServer(http.Dir(dwldPath))))

    http.ListenAndServe(":"+port, newRouter)
	
}
