package user

import "net/http"

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	
	mux.Handle("GET /user/register", http.HandlerFunc(h.CreateUser))

	mux.Handle("POST /user/profile", manager.AuthMiddleware(http.HandlerFunc(h.GetUserProfile)))
}
