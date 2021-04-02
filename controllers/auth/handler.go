package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"projet_wiki/models"
	"net/http"
	"os"
)

const (	
	JWTUsername = "username"
)

type claimsStruct struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

const JwtCookieName = "token"

func GetRequestUser(r *http.Request) (*models.User, error) {
	cookie, err := r.Cookie(JwtCookieName)
	if err != nil {
		return nil, fmt.Errorf("err : %v", err)
		
	}

	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(cookie.Value, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_KEY")), nil
	})
	if err != nil {
		return nil, fmt.Errorf("err : %v", err)
	}

	var RequestUser models.User

	for key, val := range claims {
		if key == JWTUsername {
			RequestUser.Username = fmt.Sprintf("%v", val)
		}
	}

	if RequestUser.Username == "" {
		fmt.Println("halp")
		return nil, fmt.Errorf("could not retrieve username, issue with cookie")
	}

	return &RequestUser, nil
}
