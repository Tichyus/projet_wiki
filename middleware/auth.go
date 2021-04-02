package middleware

import (
	"net/http"
	"os"
	"strings"
	"github.com/dgrijalva/jwt-go"
)

// check token -> refresh toker -> (check user)
// existence -> format du token -> (parse) -> signature -> format du payload

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var jwtKey = []byte(os.Getenv("JWT_KEY"))

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
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// We parse the token string and check signature
		claims := &Claims{}
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

		// Proceed to the controller
		next(w, r)
	}
}

// func Refresh(w http.ResponseWriter, r *http.Request) {
// 	var jwtKey = []byte(os.Getenv("JWT_KEY"))
	
// 	// (BEGIN) The code uptil this point is the same as the first part of the `Welcome` route
// 	c, err := r.Cookie("token")
// 	if err != nil {
// 		if err == http.ErrNoCookie {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			return
// 		}
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	tknStr := c.Value
// 	claims := &Claims{}
// 	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
// 		return jwtKey, nil
// 	})
// 	if !tkn.Valid {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		return
// 	}
// 	if err != nil {
// 		if err == jwt.ErrSignatureInvalid {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			return
// 		}
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
	
// 	// (END) The code uptil this point is the same as the first part of the `Welcome` route

// 	// We ensure that a new token is not issued until enough time has elapsed
// 	// In this case, a new token will only be issued if the old token is within
// 	// 30 seconds of expiry. Otherwise, return a bad request status
// 	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	// Now, create a new token for the current use, with a renewed expiration time
// 	expirationTime := time.Now().Add(5 * time.Minute)
// 	claims.ExpiresAt = expirationTime.Unix()
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	tokenString, err := token.SignedString(jwtKey)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	// Set the new token as the users `session_token` cookie
// 	http.SetCookie(w, &http.Cookie{
// 		Name:    "session_token",
// 		Value:   tokenString,
// 		Expires: expirationTime,
// 	})
// }

// func parseToken(){

// }
