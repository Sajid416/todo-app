package middlewares

import (
	"net/http"
	"strings"

	
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			SendData(w, "Missing authorization header", http.StatusUnauthorized)
			return
		}
		tokenStr := strings.TrimPrefix(tokenString, "Bearer ")
		err := VerifyToken(tokenStr)

		if err != nil {
			SendData(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
