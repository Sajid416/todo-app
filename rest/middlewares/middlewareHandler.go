package middlewares

import "github.com/Sajid416/todo-app/config"

type Middlewares struct {
	Cnf *config.Config
}

func NewMiddlewares(cnf *config.Config) *Middlewares {
	return &Middlewares{
		Cnf: cnf,
	}
}
