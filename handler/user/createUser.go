package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type ReqCreateUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}


func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ReqCreateUser
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	
	if req.Username == "" || req.Email == "" || req.Password == "" {
		http.Error(w, "Username, Email and Password are required", http.StatusBadRequest)
		return
	}

	
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}


	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id`
	var id int
	err = h.UserDB.QueryRow(query, req.Username, req.Email, string(hashedPassword)).Scan(&id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to insert user: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":  true,
		"message":  "User created successfully",
		"user_id":  id,
		"username": req.Username,
		"email":    req.Email,
	})
}
