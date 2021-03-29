package main

import (
	"log"
	"net/http"
	"projet_wiki/router"
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
