package controllers

import (
	"encoding/json"
	"net/http"

	"projet_wiki/database"
	"projet_wiki/models"

	"github.com/gorilla/mux"
)

/**
* Returns an specific user
* Required arguments: string ID
 */
func ReadUser(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn
	vars := mux.Vars(r)
	id := vars["ID"]
	var user models.User
	db.Where("ID = ?", id).Find(&user)

	json.NewEncoder(w).Encode(user)
}

/**
* Creates an user
* Required arguments: string username
 */
func CreateUser(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn
	vars := mux.Vars(r)
	username := vars["username"]

	db.Create(&models.User{Username: username, Password: "password"})
}

/**
* Deletes an user
* Required arguments: string ID
 */
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn
	vars := mux.Vars(r)
	id := vars["ID"]

	var user models.User
	db.Where("ID = ?", id).Find(&user)
	db.Delete(&user)
}

/**
* Updates an user
* Required arguments: string ID, string username
 */
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn
	vars := mux.Vars(r)
	id := vars["ID"]
	username := vars["username"]

	var user models.User
	db.Where("ID = ?", id).Find(&user)

	user.Username = username

	db.Save(&user)
}
