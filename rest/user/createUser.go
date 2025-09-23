package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Sajid416/todo-app/model"
	"github.com/Sajid416/todo-app/rest/middlewares"
	"golang.org/x/crypto/bcrypt"
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
		fmt.Println("JSON Decode error:", err)
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	// Validate input data
	if strings.TrimSpace(newUser.Username) == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(newUser.Email) == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(newUser.Password) == "" {
		http.Error(w, "Password is required", http.StatusBadRequest)
		return
	}

	// Check if user already exists
	var existingUser model.UserInfo
	checkQuery := `SELECT id FROM users WHERE email = $1 OR username = $2`
	err = h.middlewares.DB.Get(&existingUser, checkQuery, newUser.Email, newUser.Username)
	if err == nil {
		// User exists
		http.Error(w, "User with this email or username already exists", http.StatusConflict)
		return
	}

	// Hash the password
	hashPass, err := bcrypt.GenerateFromPassword(
		[]byte(strings.TrimSpace(newUser.Password)),
		bcrypt.DefaultCost,
	)
	if err != nil {
		fmt.Println("Password hashing error:", err)
		http.Error(w, "Error in password hashing", http.StatusInternalServerError)
		return
	}

	// Insert new user into database
	query := `INSERT INTO users (username, email, password)
			  VALUES ($1, $2, $3)
			  RETURNING id, username, email`

	var createdUser ReqCreateUser
	err = h.middlewares.DB.Get(&createdUser, query,
		strings.TrimSpace(newUser.Username),
		strings.TrimSpace(newUser.Email),
		string(hashPass))

	if err != nil {
		fmt.Println("Database insertion error:", err)
		http.Error(w, "Failed to create user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Send success response
	middlewares.SendData(w, createdUser, http.StatusCreated)
}
