package handler

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Sajid416/todo-app/config"
	"github.com/Sajid416/todo-app/handler/todo"
	"github.com/Sajid416/todo-app/handler/user"
	"github.com/Sajid416/todo-app/middlewares"
)

type Server struct {
	todoHandler *todo.Handler
	userHandler *user.Handler
	manager     *middlewares.Manager
}

// NewServer accepts both handlers and middleware manager
func NewServer(todoHandler *todo.Handler, userHandler *user.Handler, manager *middlewares.Manager) *Server {
	return &Server{
		todoHandler: todoHandler,
		userHandler: userHandler,
		manager:     manager,
	}
}

func (s *Server) Start(cnf *config.Config) {
	mux := http.NewServeMux()

	// ✅ Register both Todo and User routes
	s.todoHandler.RegisterRoutes(mux)
	s.userHandler.RegisterRoutes(mux)

	// ✅ Wrap mux with middlewares
	wrappedMux := s.manager.WrapMux(mux)

	addr := ":" + strconv.Itoa(cnf.HttpPort)
	log.Println("🚀 Server running on port", addr)

	if err := http.ListenAndServe(addr, wrappedMux); err != nil {
		log.Println("❌ Error starting the server:", err)
		os.Exit(1)
	}
}
