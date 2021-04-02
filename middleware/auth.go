package middleware

import (
	"net/http"
	"os"
	"strings"
	"github.com/dgrijalva/jwt-go"
)

// check token -> refresh toker -> (check user)
// existence -> format du token -> (parse) -> signature -> format du payload

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var jwtKey = []byte(os.Getenv("JWT_KEY"))

		type Claims struct {
			Username string `json:"username"`
			jwt.StandardClaims
		}

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

// func parseToken(){

// }
