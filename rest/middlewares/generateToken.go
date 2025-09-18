package middlewares

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (m *Middlewares) GenerateToken(username string, email string, expiry time.Duration,jwtToken string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"email":    email,
			"exp":      time.Now().Add(expiry).Unix(),
		})

	// FIXED: Convert secret to []byte and use HS256
	tokenString, err := claims.SignedString([]byte(jwtToken))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
