package handler

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/Sajid416/todo-app/database"
	"github.com/Sajid416/todo-app/model"
)

func SearchTaskByTitle(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")

	var user []model.User

	err := database.DB.Select(&user, `select * from tasks where title ILIKE $1`, "%"+title+"%")
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Not Found Task", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	WriteJson(w, http.StatusFound, user)
}
