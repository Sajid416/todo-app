package middlewares

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (m *Middlewares) RefreshHandler(w http.ResponseWriter, r *http.Request) {

	var req struct {
		RefreshToken string `json:"refresh_token"`
	}
	json.NewDecoder(r.Body).Decode(&req)

	token, err := jwt.Parse(req.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(m.Cnf.JWTRefresh), nil
	})

	if err != nil || !token.Valid {
		http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	userName := claims["username"].(string)
	email := claims["email"].(string)
	newAccessToken, _ := m.GenerateToken(userName, email, 15*time.Minute, m.Cnf.JWTSecret)
	

	SendData(w, map[string]string{"access_token": newAccessToken, "refresh_token": req.RefreshToken}, http.StatusOK)

}
