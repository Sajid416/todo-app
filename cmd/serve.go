package cmd

import (
	"github.com/Sajid416/todo-app/config"
	"github.com/Sajid416/todo-app/handler"
	"github.com/Sajid416/todo-app/handler/todo"
	"github.com/Sajid416/todo-app/handler/user"
	"github.com/Sajid416/todo-app/middlewares"
)

func Serve() {
	cnf := config.GetConfig()

	// Create middleware manager
	manager := middlewares.NewManager()
	manager.Use(
		middlewares.Preflight,
		middlewares.Cors,
		middlewares.Logger,
	)

	// Create handlers
	todoHandler := todo.NewHandler(manager)
	userHandler := user.NewHandler()

	// Create and start server with both handlers
	server := handler.NewServer(todoHandler, userHandler, manager)
	server.Start(cnf)
}
