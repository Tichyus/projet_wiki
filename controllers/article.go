package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"projet_wiki/database"
	"projet_wiki/models"

	"github.com/gorilla/mux"
)

/**
* Returns all articles
* Required arguments:
 */
func AllArticles(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn
	var articles []models.Article
	err := db.Find(&articles)
	if err != nil {
		fmt.Println(err)
	}

	json.NewEncoder(w).Encode(articles)
}

/**
* Returns all articles from a specific user
* Required arguments: string userId
 */
func AllArticlesFromUser(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	var articles []models.Article
	err2 := db.Where(&models.Article{UserId: id}, "User.ID").Find(&articles)
	if err2 != nil {
		fmt.Println(err)
	}

	json.NewEncoder(w).Encode(articles)
}

/**
* Returns a specific article
* Required arguments: string ID
 */
func ReadArticle(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn
	vars := mux.Vars(r)
	id := vars["id"]
	var article models.Article
	err := db.Where("ID = ?", id).Find(&article)
	if err != nil {
		fmt.Println(err)
	}

	json.NewEncoder(w).Encode(article)
}

/**
* Creates an article
* Required arguments: string title, string content, string userID
 */
func CreateArticle(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn
	title := r.FormValue("title")
	content := r.FormValue("content")
	userID, err := strconv.ParseUint(r.FormValue("userId"), 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	var user models.User
	err2 := db.First(&user, &userID)
	if err2 != nil {
		fmt.Println(err)
	}
	newArticle := &models.Article{Title: title, Content: content, User: user}

	err3 := db.Create(&newArticle)
	if err3 != nil {
		fmt.Println(err)
	}

	json.NewEncoder(w).Encode(newArticle)
}

/**
* Deletes an article
* Required arguments: ID
 */
func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn
	id := r.FormValue("id")

	var article models.Article
	err := db.Where("ID = ?", id).Find(&article)
	if err != nil {
		fmt.Println(err)
	}
	db.Delete(&article)
}

/**
* Updates an article
* Required arguments: string ID, string title, string content
 */
func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn
	id := r.FormValue("id")
	title := r.FormValue("title")
	content := r.FormValue("content")

	var article models.Article
	err := db.Where("ID = ?", id).Find(&article)
	if err != nil {
		fmt.Println(err)
	}

	article.Title = title
	article.Content = content

	err2 := db.Save(&article)
	if err2 != nil {
		fmt.Println(err)
	}

	json.NewEncoder(w).Encode(article)
}
