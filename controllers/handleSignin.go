package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"projet_wiki/models"
	"time"

	"github.com/dgrijalva/jwt-go"
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

	// DB check for user
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

	w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", tokenString))
}
