package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("your_secret_key") //!TODO check to env

func JwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		token, err := formatTokenString(tokenString)
		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func formatTokenString(rawToken string) (*jwt.Token, error) {
	tokenString := strings.TrimPrefix(rawToken, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the token method is HS256
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
		}

		return jwtSecret, nil
	})
	if err != nil {
		fmt.Println("Error while parsing JWT tokens , ", err.Error())
		return &jwt.Token{}, err
	}
	if !token.Valid {
		fmt.Println("JWT Token is invalid, Please login in again")
	}
	return token, nil
}
