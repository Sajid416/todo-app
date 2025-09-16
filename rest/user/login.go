package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Sajid416/todo-app/database"
)

type ReqLogin struct {
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

	usr := database.Find(reqLogin.Email, reqLogin.Password)
	if usr == nil {
		http.Error(w, "Invalid credentials", http.StatusBadRequest)
		return
	}

	util.SendData(w, usr, http.StatusCreated)
}
