package product

import (
	"github.com/jmoiron/sqlx"
)

type Handler struct {
	TodoDB *sqlx.DB
}

func NewHandler(TodoDB *sqlx.DB) *Handler {
	return &Handler{TodoDB: TodoDB}
}
