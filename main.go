package main

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
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
	myRouter.HandleFunc("/article/update/{id}/{title}/{content}", UpdateArticle).Methods("PUT")
	myRouter.HandleFunc("/article/delete/{id}", DeleteArticle).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	// connect to database
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
        panic("failed to connect database")
    }
    defer db.Close()

	handleRequests()
}
