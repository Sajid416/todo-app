package util

import (
	"fmt"

	"github.com/Sajid416/todo-app/config"
	"github.com/golang-jwt/jwt/v5"
)

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		cnf := config.GetConfig()
		return cnf.JWTSecret, nil
	})
	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
