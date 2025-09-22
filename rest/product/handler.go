package product

import (
	"github.com/Sajid416/todo-app/rest/middlewares"
)

type Handler struct {
	middlewares *middlewares.Middlewares
}

// TodoDB *sqlx.DB
func NewHandler(middlewares *middlewares.Middlewares) *Handler {
	return &Handler{
		middlewares: middlewares,
	}
}

//return &Handler{TodoDB: TodoDB}
