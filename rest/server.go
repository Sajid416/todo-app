package rest

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Sajid416/todo-app/config"
	"github.com/Sajid416/todo-app/rest/middlewares"
	"github.com/Sajid416/todo-app/rest/todo"
	"github.com/Sajid416/todo-app/rest/user"
)

type Server struct {
	todoHandler *todo.Handler
	userHandler *user.Handler
	manager     *middlewares.Manager
}

func NewServer(todoHandler *todo.Handler, userHandler *user.Handler, manager *middlewares.Manager) *Server {
	return &Server{
		todoHandler: todoHandler,
		userHandler: userHandler,
		manager:     manager,
	}
}

func (s *Server) Start(cnf *config.Config) {
	mux := http.NewServeMux()

	s.todoHandler.RegisterRoutes(mux, s.manager)
	s.userHandler.RegisterRoutes(mux, s.manager)

	wrappedMux := s.manager.WrapMux(mux)

	addr := ":" + strconv.Itoa(cnf.HttpPort)
	log.Println("Server running on port", addr)

	if err := http.ListenAndServe(addr, wrappedMux); err != nil {
		log.Println("Error starting the server:", err)
		os.Exit(1)
	}
}
