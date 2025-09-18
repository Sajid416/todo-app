package product

import (
	"github.com/Sajid416/todo-app/rest/middlewares"
	"github.com/jmoiron/sqlx"
)

type Handler struct {
	middlewares *middlewares.Middlewares
	ProductDB      *sqlx.DB
}

// TodoDB *sqlx.DB
func NewHandler(middlewares *middlewares.Middlewares, ProductDB *sqlx.DB) *Handler {
	return &Handler{
		middlewares: middlewares,
		ProductDB:      ProductDB,
	}
}

//return &Handler{TodoDB: TodoDB}
