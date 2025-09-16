package cmd

import (
	"log"

	"github.com/Sajid416/todo-app/config"
	"github.com/Sajid416/todo-app/rest"
	"github.com/Sajid416/todo-app/rest/middlewares"
	"github.com/Sajid416/todo-app/rest/todo"
	"github.com/Sajid416/todo-app/rest/user"
	"github.com/jmoiron/sqlx"
)

func Serve() {

	cnf := config.GetConfig()
	userDB, err := sqlx.Connect("postgres", cnf.UserDBUrl)
	if err != nil {
		log.Fatalf("Failed to connect to UserDB: %v", err)
	}
	todoDB,err:=sqlx.Connect("postgres",cnf.TodoDBUrl)
	if err!=nil{
		log.Fatalf("Failed to connect to todoDB: %v",err)
	}

	manager := middlewares.NewManager()
	manager.Use(
		middlewares.Preflight,
		middlewares.Cors,
		middlewares.Logger,
	)

	todoHandler := todo.NewHandler(todoDB)
	userHandler := user.NewHandler(userDB)
	server := rest.NewServer(todoHandler, userHandler, manager)
	server.Start(cnf)
}
