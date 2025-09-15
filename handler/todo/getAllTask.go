package handler

import (
	"net/http"

	"github.com/Sajid416/todo-app/database"
	"github.com/Sajid416/todo-app/model"
)

func GetAllTask(w http.ResponseWriter, r *http.Request) {
	var user []model.User
	rows, err := database.DB.Query(`select id,title,done,created_at from tasks order by id`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var r model.User
		if err := rows.Scan(&r.Id, &r.Title, &r.Status, &r.Created_At); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		user = append(user, r)

	}

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	WriteJson(w, http.StatusOK, user)

}
