package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/Sajid416/todo-app/database"
	"github.com/Sajid416/todo-app/model"
)

func WriteJson(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)

}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	param := strings.TrimPrefix(r.URL.Path, "/task/")
	id, err := strconv.Atoi(param)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	query := `update tasks set title=$1, status=$2 where id=$3 returning id,title,status, created_at`
	err = database.DB.Get(&user, query, user.Title, user.Status, id)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "task not found", http.StatusNotFound)
			return
		}
		http.Error(w, "failed to update task: "+err.Error(), http.StatusInternalServerError)
		return
	}

	WriteJson(w, http.StatusOK, user)
}
func UpdateStatus(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	idx := strings.TrimPrefix(r.URL.Path, "/task/")

	id, err := strconv.Atoi(idx)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	query := `UPDATE tasks SET status=$1 WHERE id=$2`
	result, err := database.DB.Exec(query, status, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	WriteJson(w, http.StatusOK, fmt.Sprintf("Update successful, %d row(s) affected", rowsAffected))
}
