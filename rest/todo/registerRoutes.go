package todo

import (
	"net/http"

	"github.com/Sajid416/todo-app/rest/middlewares"
)

func RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager){

	mux.Handle("GET /task", manager.WrapHandler(
	http.HandlerFunc(GetAllTask),
))

mux.Handle("GET /task/filter", manager.WrapHandler(
	http.HandlerFunc(FilteredTask),
	
))

	mux.Handle("POST /task",manager.WrapHandler(http.HandlerFunc(CreateTask)))
	mux.Handle("PUT /task/{id}",manager.WrapHandler(http.HandlerFunc(UpdateTask)))
	mux.Handle("DELETE /{id}",manager.WrapHandler(http.HandlerFunc(DeleteTask)))
	mux.Handle("GET /task",manager.WrapHandler(http.HandlerFunc(FilteredTask)))
}