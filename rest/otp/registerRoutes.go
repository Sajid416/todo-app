package otp

import (
	"net/http"

	"github.com/Sajid416/todo-app/rest/middlewares"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle("/send-otp", manager.With(http.HandlerFunc(h.SendOTP)))
	mux.Handle("/verify-otp", manager.With(http.HandlerFunc(h.VerifyOTP)))
}
