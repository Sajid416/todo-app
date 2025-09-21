package product

import (
	"github.com/Sajid416/todo-app/rest/middlewares"
	"github.com/jmoiron/sqlx"
)

type Handler struct {
	middlewares *middlewares.Middlewares
	DBUrl       *sqlx.DB
}

// TodoDB *sqlx.DB
func NewHandler(middlewares *middlewares.Middlewares, DBUrl *sqlx.DB) *Handler {
	return &Handler{
		middlewares: middlewares,
		DBUrl:       DBUrl,
	}
}

//return &Handler{TodoDB: TodoDB}
