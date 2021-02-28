package controllers

import (
	"net/http"
	"gorm.io/gorm"
	"github.com/gorilla/mux"
)

func AllArticles(w http.ResponseWriter, r *http.Request) {
	var articles []Article
    db.Find(&articles)

    json.NewEncoder(w).Encode(articles)
}

func AllArticlesFromUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
    var articles []Article
	var article []Article
	db.Where(&Article{author.id: id}, "author.id").Find(&articles)

    json.NewEncoder(w).Encode(articles)
}

func ReadArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
    var article Article
    db.Where("id = ?", id).Find(&article)

    json.NewEncoder(w).Encode(article)
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    title := vars["title"]
    content := vars["content"]
	user := vars["user"]

    db.Create(&Article{Title: title, Content: content})
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    id := vars["id"]

    var article Article
    db.Where("id = ?", id).Find(&article)
    db.Delete(&article)
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
    title := vars["title"]
    content := vars["content"]

    var article Article
    db.Where("id = ?", id).Find(&article)

	article.Title = title
    article.Content = content

    db.Save(&article)
}
