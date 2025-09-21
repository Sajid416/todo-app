package user

import (
	"net/http"

	"github.com/Sajid416/todo-app/rest/middlewares"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager) {

	mux.Handle("POST /user/register", manager.With(http.HandlerFunc(h.CreateUser)))

	mux.Handle("POST /user/login", manager.With(http.HandlerFunc(h.Login)))
}
