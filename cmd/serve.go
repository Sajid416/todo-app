package cmd

import (
	"github.com/Sajid416/todo-app/handler/user"
	"github.com/Sajid416/todo-app/middlewares"
)

func Serve(){
	cnf:=config.GetConfig()
	middlewares:=middlewares.NewMiddlewares(cnf)
	productHandler:=product.NewHandler(middlewares)
	userHandler:=user.NewHandler()
	reviewHandler:=review.NewHandler()

	server:=rest.NewServer(
		cnf,
		productHandler,
		userHandler,
		reviewHandler,
	)
	server.Start()
}