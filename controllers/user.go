package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"projet_wiki/database"
	"projet_wiki/models"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

/**
* Returns an specific user
* Required arguments: string ID
 */
func ReadUser(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn
	vars := mux.Vars(r)
	ID := vars["ID"]
	var user models.User
	err := db.Where("id = ?", ID).Find(&user)
	if err != nil {
		fmt.Println(err)
	}

	json.NewEncoder(w).Encode(user)
}

/**
* Creates an user
* Required arguments: string username
 */
func CreateUser(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn
	username := r.FormValue("username")
	password := r.FormValue("password")
	bytes, err1 := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err1 != nil {
		fmt.Println(err1)
	}
	hashedPassword := string(bytes)

	user := &models.User{Username: username, Password: hashedPassword}

	err2 := db.Create(&user)
	if err2 != nil {
		fmt.Println(err2)
	}

	json.NewEncoder(w).Encode(user)
}

/**
* Deletes an user
* Required arguments: string ID
 */
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn
	ID := r.FormValue("ID")

	var user models.User
	err := db.Where("id = ?", ID).Find(&user)
	if err != nil {
		fmt.Println(err)
	}
	db.Delete(&user)
}

/**
* Updates an user
* Required arguments: string ID, string username
 */
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn
	ID := r.FormValue("ID")
	username := r.FormValue("username")

	var user models.User
	err := db.Where("id = ?", ID).Find(&user)
	if err != nil {
		fmt.Println(err)
	}

	user.Username = username

	err2 := db.Save(&user)
	if err2 != nil {
		fmt.Println(err)
	}

	json.NewEncoder(w).Encode(user)
}

func CheckUserAuthCreds(username string, password string) bool {
	db := database.DbConn
	var user models.User
	err := db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return false
	}
	if password != user.Password {
		return false
	}
	return true
}
