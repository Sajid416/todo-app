package cmd

import (
	"github.com/Sajid416/todo-app/config"
	"github.com/Sajid416/todo-app/rest"
	"github.com/Sajid416/todo-app/rest/product"
	"github.com/Sajid416/todo-app/rest/user"
)

func Serve() {

	cnf := config.GetConfig()
	// userDB, err := sqlx.Connect("postgres", cnf.UserDBUrl)
	// if err != nil {
	// 	log.Fatalf("Failed to connect to UserDB: %v", err)
	// }
	// productDB, err := sqlx.Connect("postgres", cnf.TodoDBUrl)
	// if err != nil {
	// 	log.Fatalf("Failed to connect to productDB: %v", err)
	// }

	// manager := middlewares.NewManager()
	// manager.Use(
	// 	middlewares.Preflight,
	// 	middlewares.Cors,
	// 	middlewares.Logger,
	// )

	// productHandler := product.NewHandler()
	// userHandler := user.NewHandler()
	// server := rest.NewServer(productHandler, userHandler, manager)
	productHandler := product.NewHandler()
	userHandler := user.NewHandler()

	server := rest.NewServer(cnf, productHandler, userHandler)
	server.Start()
}
