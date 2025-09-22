package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Sajid416/todo-app/model"
	"github.com/Sajid416/todo-app/rest/middlewares"
)

type ReqLogin struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var reqLogin ReqLogin

	if err := json.NewDecoder(r.Body).Decode(&reqLogin); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		fmt.Println("JSON Decode error:", err)
		return
	}
	var user model.UserInfo
	query := `select * from users where email=$1`
	err := h.middlewares.DB.Get(&user, query, reqLogin.Email)
	if err != nil {
		middlewares.SendData(w, "Invalid Email or Password", http.StatusUnauthorized)
		return
	}
	fmt.Println(len(user.Password))
	reqLogin.Password = strings.TrimSpace(reqLogin.Password)
	if !middlewares.Compare_Pass(reqLogin.Password, user.Password) {
		middlewares.SendData(w, "invalid password or email", http.StatusUnauthorized)
		return
	}

	accessToken, err := h.middlewares.GenerateToken(reqLogin.Username, reqLogin.Email, 15*time.Minute, h.middlewares.Cnf.JWTSecret)
	if err != nil {
		middlewares.SendData(w, "Failed to create access token", http.StatusInternalServerError)
		return
	}
	refreshToken, err := h.middlewares.GenerateToken(reqLogin.Username, reqLogin.Email, 7*24*time.Hour, h.middlewares.Cnf.JWTRefresh)
	if err != nil {
		middlewares.SendData(w, "Failed to refresh token", http.StatusInternalServerError)
		return
	}
	middlewares.SendData(w, map[string]string{"access_token": accessToken, "refresh_token": refreshToken}, http.StatusOK)
}
