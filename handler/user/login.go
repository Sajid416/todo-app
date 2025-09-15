package user

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Sajid416/todo-app/model"
	"golang.org/x/crypto/bcrypt"
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

	var usr model.UserInfo
	query := `SELECT id, email, password FROM users WHERE email=$1`
	err := h.UserDB.Get(&usr, query, reqLogin.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			fmt.Println("DB error:", err)
		}
		return
	}


	if err := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(reqLogin.Password)); err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := map[string]interface{}{
		"message": "Login successful",
		"user_id": usr.ID,
		"email":   usr.Email,
	}
	json.NewEncoder(w).Encode(resp)
}
