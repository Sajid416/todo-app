package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Sajid416/todo-app/model"
	"github.com/Sajid416/todo-app/rest/middlewares"
)

type ReqLogin struct {
	ID       int    `db:"id"`
	UserName string `json:"userName"`
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
	query := `select * from user where email=$1`
	err := h.UserDB.Get(&user, query, reqLogin.Email)
	if err != nil {
		middlewares.SendData(w, "Invalid Email or Password", http.StatusUnauthorized)
		return
	}
	if !middlewares.Compare_Pass(reqLogin.Password, user.Password) {
		middlewares.SendData(w, "invalid password or email", http.StatusUnauthorized)
		return
	}

	token, err := h.Middlewares.GenerateToken(reqLogin.UserName, reqLogin.Email)
    if err!=nil{
		middlewares.SendData(w,"Failed to create token",http.StatusInternalServerError)
		return 
	}
	middlewares.SendData(w, map[string]string{"token":token}, http.StatusOK)
}
