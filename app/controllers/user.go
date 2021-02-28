package controllers

import (
	"gorm.io/gorm"
	"github.com/gorilla/mux"
)

func ReadUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
	id := vars["id"]
    var user User
    db.Where("id = ?", id).Find(&user)

    json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    username := vars["username"]

    db.Create(&User{Username: username})
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    id := vars["id"]

    var user User
    db.Where("id = ?", id).Find(&user)
    db.Delete(&user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
    username := vars["username"]

    var user User
    db.Where("id = ?", id).Find(&user)

    user.Username = username

    db.Save(&user)
}
