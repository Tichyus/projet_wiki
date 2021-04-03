package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"projet_wiki/controllers/auth"
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
	ID := vars["ID"]
	var comment models.Comment
	err := db.Where("id= ?", ID).Find(&comment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(comment)
}

/**
* Returns all comments from a specific article
* Required arguments: string ID
 */
func ReadComments(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn
	vars := mux.Vars(r)
	ID, err := strconv.ParseUint(vars["ID"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var comments []models.Comment
	err2 := db.Where("article_id = ?", ID).Find(&comments)
	if err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(comments)
}

/**
* Creates a comment
* Required arguments: string content, string userID, string articleID
 */
func CreateComment(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn
	content := r.FormValue("content")
	userID, err := strconv.ParseUint(r.FormValue("userID"), 10, 32)
	articleID, err := strconv.ParseUint(r.FormValue("articleID"), 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var user models.User
	err2 := db.First(&user, &userID)
	if err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var article models.Article
	err3 := db.First(&article, &articleID)
	if err3 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	newComment := &models.Comment{Content: content, Article: article, User: user}

	err4 := db.Create(&newComment)
	if err4 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newComment)
}

/**
* Deletes a comment
* Required arguments: string ID
 */
func DeleteComment(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn
	ID := r.FormValue("ID")

	var comment models.Comment
	err := db.Where("id = ?", ID).Find(&comment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	UserAuth, errAuth := auth.GetRequestUser(r)
	if errAuth != nil {
		fmt.Println("user not connected")
		return
	}

	UserId := comment.UserID
	var user models.User
	db.Where("id = ?", UserId).Find(&user)
	
	if err != nil {
		fmt.Println(err)
	}

	if UserAuth.Username != user.Username {
		fmt.Println("Incorrect account rights ")
	    return
	}

	db.Delete(&comment)
}

/**
* Updates a comment
* Required arguments: string ID, string content
 */
func UpdateComment(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn
	ID := r.FormValue("ID")
	content := r.FormValue("content")

	var comment models.Comment
	err := db.Where("id = ?", ID).Find(&comment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	UserAuth, errAuth := auth.GetRequestUser(r)
	if errAuth != nil {
		fmt.Println("user not connected")
		return
	}

	UserId := comment.UserID
	var user models.User
	db.Where("id = ?", UserId).Find(&user)
	
	if err != nil {
		fmt.Println(err)
	}

	if UserAuth.Username != user.Username {
		fmt.Println("Incorrect account rights ")
	    return
	}

	comment.Content = content

	err2 := db.Save(&comment)
	if err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(comment)
}
