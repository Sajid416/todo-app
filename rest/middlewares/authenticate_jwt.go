package middlewares

import (
	"net/http"
)

func (m *Manager) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Example: Extract token from Authorization header
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		// TODO: Verify JWT properly (using m.JWTSecret)
		// For now, just simulate check
		if token != "Bearer VALID_TOKEN" {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Call next handler if token is valid
		next.ServeHTTP(w, r)
	})
}
