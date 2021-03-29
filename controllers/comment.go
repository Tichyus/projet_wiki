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
* Returns a specific comment
* Required arguments: string ID
 */
func ReadComment(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn
	vars := mux.Vars(r)
	id := vars["ID"]
	var comment models.Comment
	db.Where("ID = ?", id).Find(&comment)

	json.NewEncoder(w).Encode(comment)
}

/**
* Returns all comments from a specific article
* Required arguments: string ID
 */
func ReadComments(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["ID"], 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	var article models.Article

	comments := db.Where("ID = ?", id).Find(&article).Association("Comments")

	json.NewEncoder(w).Encode(comments)
}

/**
* Creates a comment
* Required arguments: string content, string userID, string articleID
 */
func CreateComment(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn
	vars := mux.Vars(r)
	content := vars["content"]
	userID, err := strconv.ParseUint(vars["userID"], 10, 32)
	articleID, err := strconv.ParseUint(vars["articleID"], 10, 32)
	if err != nil {
		fmt.Println(err)
	}

	db.Create(&models.Comment{Content: content, ArticleId: articleID, UserId: userID})
}

/**
* Deletes a comment
* Required arguments: string ID
 */
func DeleteComment(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn
	vars := mux.Vars(r)
	id := vars["ID"]

	var comment models.Comment
	db.Where("ID = ?", id).Find(&comment)
	db.Delete(&comment)
}

/**
* Updates a comment
* Required arguments: string ID, string content
 */
func UpdateComment(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn
	vars := mux.Vars(r)
	id := vars["ID"]
	content := vars["content"]

	var comment models.Comment
	db.Where("ID = ?", id).Find(&comment)

	comment.Content = content

	db.Save(&comment)
}
