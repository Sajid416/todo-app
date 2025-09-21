package user

import (
	"github.com/Sajid416/todo-app/rest/middlewares"
	"github.com/jmoiron/sqlx"
)

type Handler struct {
	Middlewares *middlewares.Middlewares
	DBUrl *sqlx.DB
}

// userDB *sqlx.DB
func NewHandler(Middlewares *middlewares.Middlewares,DBUrl *sqlx.DB) *Handler {
	return &Handler{
		Middlewares:Middlewares,
		DBUrl: DBUrl,
	}
}

//return &Handler{UserDB: userDB}
