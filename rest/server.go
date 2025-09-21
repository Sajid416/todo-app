package rest

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Sajid416/todo-app/config"
	"github.com/Sajid416/todo-app/rest/middlewares"
	"github.com/Sajid416/todo-app/rest/otp"
	"github.com/Sajid416/todo-app/rest/product"
	"github.com/Sajid416/todo-app/rest/user"
)

type Server struct {
	cnf            *config.Config
	productHandler *product.Handler
	userHandler    *user.Handler
	otpHandler *otp.Handler
}

func NewServer(cnf *config.Config, productHandler *product.Handler, userHandler *user.Handler, otpHandler *otp.Handler) *Server {
	return &Server{
		cnf:            cnf,
		productHandler: productHandler,
		userHandler:    userHandler,
		otpHandler: otpHandler,

	}
}

func (server *Server) Start() {
	manager := middlewares.NewManager()
	manager.Use(
		middlewares.Preflight,
		middlewares.Cors,
		middlewares.Logger,
	)

	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)
	server.otpHandler.RegisterRoutes(mux,manager)
	server.productHandler.RegisterRoutes(mux, manager)
	server.userHandler.RegisterRoutes(mux, manager)

	addr := ":" + strconv.Itoa(server.cnf.HttpPort)
	log.Println("Server running on port", addr)

	if err := http.ListenAndServe(addr, wrappedMux); err != nil {
		log.Println("Error starting the server:", err)
		os.Exit(1)
	}
}
