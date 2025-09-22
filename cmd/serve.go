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
	DB, err := sqlx.Connect("postgres", cnf.DBUrl)
	if err != nil {
		log.Fatalf("Failed to connect to UserDB: %v", err)
	}
	rdb := otp.NewClient("localhost:6379")
	// manager := middlewares.NewManager()
	// manager.Use(
	// 	middlewares.Preflight,
	// 	middlewares.Cors,
	// 	middlewares.Logger,
	// )
	m := middlewares.NewMiddlewares(cnf,DB)
	productHandler := product.NewHandler(m)
	userHandler := user.NewHandler(m)
	manager := otp.NewManager(rdb)
	otpHandler := otp.NewHandler(manager)

	server := rest.NewServer(
		cnf,
		productHandler,
		userHandler,
		otpHandler,
	)
	server.Start()
}
