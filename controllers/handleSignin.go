package controllers

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"time"
	"projet_wiki/models"
)


func Signin(w http.ResponseWriter, r *http.Request) {
	var jwtKey = []byte(os.Getenv("JWT_KEY"))

	type Claims struct {
		Username string `json:"username"`
		jwt.StandardClaims
	}
	
	var user models.User

	// Get the JSON body and decode into credentials
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !CheckUserAuthCreds(user.Username, user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// TODO : LE METTRE EN HEADER
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}
