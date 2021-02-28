package main

import "net/http"
import "projet_wiki/Controllers/user.go"


func main() {
	fs := http.FileServer(http.Dir("./views"))
	http.Handle("/*", fs)

	http.HandleFunc("/", spaHandler);

	http.ListenAndServe(":8080", nil)

	Controllers.UserHandler()

}


func spaHandler(responseWriter http.ResponseWriter, request *http.Request) {
	http.ServeFile(responseWriter, request, "./views")
}

// package main

// import (
// 	"gorm.io/gorm"
// 	"gorm.io/driver/mysql"
// )

// func main() {
// 	// connect to database
// 	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// }