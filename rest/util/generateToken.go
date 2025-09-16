package util

import (
	"time"

	"github.com/Sajid416/todo-app/config"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(username string, email string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodES256,
		jwt.MapClaims{
			"username": username,
			"email":    email,
			"exp":      time.Now().Add(24 * time.Hour).Unix(),
		})
	cnf := config.GetConfig()
	tokenString, err := claims.SignedString(cnf.JWTSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
