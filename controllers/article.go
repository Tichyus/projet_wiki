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

func AllArticles(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn
	var articles []models.Article
	db.Find(&articles)

	json.NewEncoder(w).Encode(articles)
}

func AllArticlesFromUser(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn
	vars := mux.Vars(r)
	tempID, err := strconv.ParseUint(vars["userId"], 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	ID := uint(tempID)
	var articles []models.Article
	db.Where(&models.Article{UserId: ID}, "User.ID").Find(&articles)

	json.NewEncoder(w).Encode(articles)
}

func ReadComments(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn
	vars := mux.Vars(r)
	tempID, err := strconv.ParseUint(vars["ID"], 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	ID := uint(tempID)
	var article models.Article

	comments := db.Where("ID = ?", ID).Find(&article).Association("Comments")

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
	TempUserID, err := strconv.ParseUint(vars["userID"], 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	UserID := uint(TempUserID)

	db.Create(&models.Article{Title: title, Content: content, UserId: UserID})
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
