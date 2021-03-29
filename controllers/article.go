package controllers

import (
	"encoding/json"
	"net/http"

	"projet_wiki/database"
	"projet_wiki/models"

	"github.com/gorilla/mux"
)

func AllArticles(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn
	var articles []models.Article
	db.Find(&articles)

	json.NewEncoder(w).Encode(articles)
}

func AllArticlesFromUser(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn
	vars := mux.Vars(r)
	ID := vars["ID"]
	var articles []models.Article
	db.Where(&models.Article{author.ID: ID}, "author.ID").Find(&articles)

	json.NewEncoder(w).Encode(articles)
}

func ReadComments(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn
	vars := mux.Vars(r)
	ID := vars["ID"]
	var article models.Article

	comments := db.Where(&models.Article{ID: ID}, "ID").Find(&article).Association("Comments").Find(models.Comment)

	json.NewEncoder(w).Encode(comments)
}

func ReadArticle(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn
	vars := mux.Vars(r)
	ID := vars["ID"]
	var article models.Article
	db.Where("ID = ?", ID).Find(&article)

	json.NewEncoder(w).Encode(article)
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn
	vars := mux.Vars(r)
	title := vars["title"]
	content := vars["content"]
	user := vars["user"]

	db.Create(&models.Article{Title: title, Content: content, User: user})
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn
	vars := mux.Vars(r)
	ID := vars["ID"]

	var article models.Article
	db.Where("ID = ?", ID).Find(&article)
	db.Delete(&article)
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn
	vars := mux.Vars(r)
	ID := vars["ID"]
	title := vars["title"]
	content := vars["content"]

	var article models.Article
	db.Where("ID = ?", ID).Find(&article)

	article.Title = title
	article.Content = content

	db.Save(&article)
}
