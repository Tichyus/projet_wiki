package controllers

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"github.com/gorilla/mux"
)

func AllCommentsFromArticle(w http.ResponseWriter, r *http.Request) {
    db, err := gorm.Open("mysql", "flamingo.db")
    if err != nil {
        panic("failed to connect database")
    }
    defer db.Close()

	vars := mux.Vars(r)
	id := vars["id"]
    var comments []Comment
	var comment []Comment
	db.Where(&Comment{article.id: id}, "article.id").Find(&comments)

    json.NewEncoder(w).Encode(comments)
}

func ReadComment(w http.ResponseWriter, r *http.Request) {
    db, err := gorm.Open("mysql", "flamingo.db")
    if err != nil {
        panic("failed to connect database")
    }
    defer db.Close()

    vars := mux.Vars(r)
	id := vars["id"]
    var comment Comment
    db.Where("id = ?", id).Find(&comment)

    json.NewEncoder(w).Encode(comment)
}

func CreateComment(w http.ResponseWriter, r *http.Request) {
    db, err := gorm.Open("mysql", "flamingo.db")
    if err != nil {
        panic("failed to connect database")
    }
    defer db.Close()

    vars := mux.Vars(r)
    content := vars["content"]
	user := vars["user"]

    db.Create(&Comment{Content: content})
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {
    db, err := gorm.Open("mysql", "flamingo.db")
    if err != nil {
        panic("failed to connect database")
    }
    defer db.Close()

    vars := mux.Vars(r)
    id := vars["id"]

    var comment Comment
    db.Where("id = ?", id).Find(&comment)
    db.Delete(&comment)
}

func UpdateComment(w http.ResponseWriter, r *http.Request) {
    db, err := gorm.Open("sqlite3", "test.db")
    if err != nil {
        panic("failed to connect database")
    }
    defer db.Close()

    vars := mux.Vars(r)
	id := vars["id"]
    content := vars["content"]

    var comment Comment
    db.Where("id = ?", id).Find(&comment)

    comment.Content = content

    db.Save(&comment)
}
