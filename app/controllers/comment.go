package controllers

import (
	"net/http"
	"gorm.io/gorm"
	"github.com/gorilla/mux"
)

func ReadComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["ID"]
    var comment Comment
    db.Where("ID = ?", ID).Find(&comment)

    json.NewEncoder(w).Encode(comment)
}

func CreateComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    content := vars["content"]
	user_ID := vars["user"]
	article_ID := vars["user"]

	var user User
	var article Article
    db.Where("ID = ?", ID).Find(&comment)
    db.Create(&Comment{Content: content})
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    ID := vars["ID"]

    var comment Comment
    db.Where("ID = ?", ID).Find(&comment)
    db.Delete(&comment)
}

func UpdateComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["ID"]
    content := vars["content"]

    var comment Comment
    db.Where("ID = ?", ID).Find(&comment)

    comment.Content = content

    db.Save(&comment)
}
