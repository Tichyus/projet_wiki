package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func Signin(w http.ResponseWriter, r *http.Request) {
	var jwtKey = []byte(os.Getenv("JWT_KEY"))

	type claimsStruct struct {
		Username string `json:"username"`
		jwt.StandardClaims
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	// DB check for user
	if !CheckUserAuthCreds(username, password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(10 * time.Minute)

	claims := &claimsStruct{
		Username: username,
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
