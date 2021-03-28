package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func ReadUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["ID"]
	var user User
	db.Where("ID = ?", ID).Find(&user)

	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	db.Create(&User{Username: username})
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["ID"]

	var user User
	db.Where("ID = ?", ID).Find(&user)
	db.Delete(&user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["ID"]
	username := vars["username"]

	var user User
	db.Where("ID = ?", ID).Find(&user)

	user.Username = username

	db.Save(&user)
}
