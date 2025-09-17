package product

import (
	"github.com/Sajid416/todo-app/rest/middlewares"
	"github.com/jmoiron/sqlx"
)

type Handler struct {
	middlewares *middlewares.Middlewares
	TodoDB      *sqlx.DB
}

// TodoDB *sqlx.DB
func NewHandler(middlewares *middlewares.Middlewares, TodoDB *sqlx.DB) *Handler {
	return &Handler{
		middlewares: middlewares,
		TodoDB:      TodoDB,
	}
}

//return &Handler{TodoDB: TodoDB}
