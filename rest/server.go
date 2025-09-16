package rest

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Sajid416/todo-app/config"
	"github.com/Sajid416/todo-app/rest/middlewares"
	"github.com/Sajid416/todo-app/rest/product"
	"github.com/Sajid416/todo-app/rest/user"
)

type Server struct {
	productHandler *product.Handler
	userHandler    *user.Handler
	manager        *middlewares.Manager
}

func NewServer(productHandler *product.Handler, userHandler *user.Handler, manager *middlewares.Manager) *Server {
	return &Server{
		productHandler: productHandler,
		userHandler:    userHandler,
		manager:        manager,
	}
}

func (s *Server) Start(cnf *config.Config) {
	mux := http.NewServeMux()

	s.productHandler.RegisterRoutes(mux, s.manager)
	s.userHandler.RegisterRoutes(mux, s.manager)

	wrappedMux := s.manager.WrapMux(mux)

	addr := ":" + strconv.Itoa(cnf.HttpPort)
	log.Println("Server running on port", addr)

	if err := http.ListenAndServe(addr, wrappedMux); err != nil {
		log.Println("Error starting the server:", err)
		os.Exit(1)
	}
}
