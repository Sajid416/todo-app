package handler

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/Sajid416/todo-app/database"
	"github.com/Sajid416/todo-app/model"
)

func FilteredTask(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	query := `select * from tasks where status=$1`
	var user model.User
	err := database.DB.Select(&user, query, status)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Pending Task Not Found", http.StatusNotFound)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	WriteJson(w, http.StatusFound, "Available pending Task")
}
