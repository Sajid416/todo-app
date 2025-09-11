package router

import (
	"net/http"

	"github.com/Sajid416/todo-app/handler"
)

func RegisterRoutes() *http.ServeMux {

	r := http.NewServeMux()
	r.HandleFunc("/task", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case "GET":
			if r.URL.Query().Has("status") {
				handler.FilteredTask(w, r)
			} else if r.URL.Query().Has("title") {
				handler.UpdateStatus(w, r)

			} else {
				handler.GetAllTask(w, r)
			}
		case "POST":
			handler.CreateTask(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}

	})
	r.HandleFunc("/task/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			handler.GetTaskById(w, r)
		case "PUT":
			if r.URL.Query().Has("status") {
				handler.UpdateStatus(w, r)
			} else {
				handler.UpdateTask(w, r)
			}

		case "DELETE":
			handler.DeleteTask(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	return r

}
