package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"projet_wiki/database"
	"projet_wiki/router"
)

const dwldPath = "./tmp"

func main() {
	port := "8080"
	newRouter := router.NewRouter()

	err := database.Connect()
	if err != nil {
		log.Fatalf("Impossible de se connecter à la bdd: %v", err)
	}

	// Loading environment variables
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	log.Print("\nServer started on port " + port)

	newRouter.PathPrefix("/files/").Handler(http.StripPrefix("/files/",
		http.FileServer(http.Dir(dwldPath))))

	http.ListenAndServe(":"+port, newRouter)

}
