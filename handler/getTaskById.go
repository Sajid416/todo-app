package handler

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/Sajid416/todo-app/database"
	"github.com/Sajid416/todo-app/model"
)

func GetTaskById(w http.ResponseWriter, r *http.Request) {
	idr := strings.TrimPrefix(r.URL.Path, "/task/")
	id, err := strconv.Atoi(idr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var user model.User
	err = database.DB.Get(&user, `select * from tasks where id=$1`, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	WriteJson(w, http.StatusOK, user)

}
