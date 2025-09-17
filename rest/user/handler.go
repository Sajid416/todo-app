package user

import (
	"github.com/Sajid416/todo-app/rest/middlewares"
	"github.com/jmoiron/sqlx"
)

type Handler struct {
	Middlewares *middlewares.Middlewares
	UserDB *sqlx.DB
}

// userDB *sqlx.DB
func NewHandler(Middlewares *middlewares.Middlewares,UserDB *sqlx.DB) *Handler {
	return &Handler{
		Middlewares:Middlewares,
		UserDB: UserDB,
	}
}

//return &Handler{UserDB: userDB}
