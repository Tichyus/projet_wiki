package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func AllArticles(w http.ResponseWriter, r *http.Request) {
	var articles []Article
	db.Find(&articles)

	json.NewEncoder(w).Encode(articles)
}

func AllArticlesFromUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["ID"]
	var articles []Article
	var article []Article
	db.Where(&Article{author.ID: ID}, "author.ID").Find(&articles)

	json.NewEncoder(w).Encode(articles)
}

func ReadComments(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["ID"]
	var articles []Article
	var article []Article

	comments := db.Where(&Article{ID: ID}, "ID").Find(&article).Association("Comments").Find(&comments)

	json.NewEncoder(w).Encode(comments)
}

func ReadArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["ID"]
	var article Article
	db.Where("ID = ?", ID).Find(&article)

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
	ID := vars["ID"]

	var article Article
	db.Where("ID = ?", ID).Find(&article)
	db.Delete(&article)
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["ID"]
	title := vars["title"]
	content := vars["content"]

	var article Article
	db.Where("ID = ?", ID).Find(&article)

	article.Title = title
	article.Content = content

	db.Save(&article)
}
