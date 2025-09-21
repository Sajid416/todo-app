package cmd

import (
	"log"

	"github.com/Sajid416/todo-app/config"
	"github.com/Sajid416/todo-app/rest"
	"github.com/Sajid416/todo-app/rest/middlewares"
	"github.com/Sajid416/todo-app/rest/otp"
	"github.com/Sajid416/todo-app/rest/product"
	"github.com/Sajid416/todo-app/rest/user"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Serve() {

	cnf := config.GetConfig()
	DBUrl, err := sqlx.Connect("postgres", cnf.DBUrl)
	if err != nil {
		log.Fatalf("Failed to connect to UserDB: %v", err)
	}
    rdb:=otp.NewClient("localhost:6379")
	// manager := middlewares.NewManager()
	// manager.Use(
	// 	middlewares.Preflight,
	// 	middlewares.Cors,
	// 	middlewares.Logger,
	// )

	// productHandler := product.NewHandler()
	// userHandler := user.NewHandler()
	// server := rest.NewServer(productHandler, userHandler, manager)
	m := middlewares.NewMiddlewares(cnf)
	productHandler := product.NewHandler(m, DBUrl)
	userHandler := user.NewHandler(m, DBUrl)
    manager:=otp.NewManager(rdb)

	server := rest.NewServer(
		cnf,
		productHandler,
		userHandler,
		manager,
	)
	server.Start()
}
