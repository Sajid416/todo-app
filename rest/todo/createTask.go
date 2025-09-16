package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Sajid416/todo-app/database"
	"github.com/Sajid416/todo-app/model"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {

	var user model.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	query := `INSERT INTO tasks (title, status) VALUES ($1, $2) RETURNING id, created_at`
	err := database.DB.Get(&user, query, user.Title, user.Status)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if user.Created_At.IsZero() {
		user.Created_At = time.Now()
	}
	WriteJson(w, http.StatusCreated, user)

}
