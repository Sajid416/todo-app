package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Sajid416/todo-app/model"
	"github.com/Sajid416/todo-app/rest/middlewares"
)

type ReqCreateUser struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
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

	hashPass, err := middlewares.HashedPassword(newUser.Password)
	if err != nil {
		http.Error(w, "Error in Password hashing", http.StatusInternalServerError)
		return
	}
	query := `insert into users (username,email,password)
	        values ($1,$2,$3)
			returning id,username,email`

	var User ReqCreateUser
	err= h.UserDB.Get(&User,query,newUser.Username,newUser.Email,hashPass)
	if err!=nil{
		http.Error(w,"Failed to insert User:"+err.Error(),http.StatusInternalServerError)
	}

	middlewares.SendData(w,User, http.StatusCreated)
}
