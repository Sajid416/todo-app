package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Sajid416/todo-app/model"
	"github.com/Sajid416/todo-app/rest/middlewares"
)

type ReqCreateUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser model.UserInfo
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	middlewares.SendData(w,newUser, http.StatusCreated)
}
