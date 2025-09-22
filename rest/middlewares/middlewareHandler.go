package middlewares

import (
	"github.com/Sajid416/todo-app/config"
	"github.com/jmoiron/sqlx"
)

type Middlewares struct {
	Cnf *config.Config
	DB  *sqlx.DB
}

func NewMiddlewares(cnf *config.Config, DB *sqlx.DB) *Middlewares {
	return &Middlewares{
		Cnf: cnf,
		DB:DB,
	}
}
