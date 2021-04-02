package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type claimsStruct struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func RefreshToken(w http.ResponseWriter, r *http.Request) {
	var jwtKey = []byte(os.Getenv("JWT_KEY"))
	claims := &claimsStruct{}

	// We check token existence in the header
	tknStr := r.Header.Get("Authorization")
	if tknStr == "" {
		w.WriteHeader(http.StatusUnauthorized) // Http 401
		return
	}

	// After extraction and split, it must give an array of len = 2
	extractedToken := strings.Split(tknStr, "Bearer ")
	if len(extractedToken) == 2 {
		tknStr = strings.TrimSpace(extractedToken[1])
	} else {
		fmt.Printf("bad request 1")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	// We parse the token string and check signature
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	// new token issued with a renewed expiration time
	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", tokenString))
} 