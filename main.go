package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/user/create/{username}", CreateUser).Methods("POST")
	myRouter.HandleFunc("/user/{id}", ReadUser).Methods("GET")
	myRouter.HandleFunc("/user/update/{id}/{username}", UpdateUser).Methods("PUT")
	myRouter.HandleFunc("/user/delete/{id}", DeleteUser).Methods("DELETE")

	myRouter.HandleFunc("/comment/create/{content}", CreateComment).Methods("POST")
	myRouter.HandleFunc("/comment/{id}", ReadComment).Methods("GET")
	myRouter.HandleFunc("/comment/update/{id}/{content}", UpdateComment).Methods("PUT")
	myRouter.HandleFunc("/comment/delete/{id}", DeleteComment).Methods("DELETE")

	myRouter.HandleFunc("/article/create/{title}/{content}", CreateArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", ReadArticle).Methods("GET")
	myRouter.HandleFunc("/articles", AllArticles).Methods("GET")
	myRouter.HandleFunc("/articles/{id}", AllArticlesFromUser).Methods("GET")
	myRouter.HandleFunc("/article/{id}/comments", ReadComments).Methods("GET")
	myRouter.HandleFunc("/article/update/{id}/{title}/{content}", UpdateArticle).Methods("PUT")
	myRouter.HandleFunc("/article/delete/{id}", DeleteArticle).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	handleRequests()
}
