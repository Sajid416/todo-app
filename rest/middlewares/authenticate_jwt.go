package middlewares

import (
	"context"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			SendData(w, "Missing authorization header", http.StatusUnauthorized)
			return
		}

		// Validate "Bearer <token>" format
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			SendData(w, "Invalid authorization header format. Expected: Bearer <token>", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]
		if tokenString == "" {
			SendData(w, "Empty token", http.StatusUnauthorized)
			return
		}

		// Verify and parse token
		userClaims, err := VerifyToken(tokenString)
		if err != nil {
			SendData(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Add user information to request context
		ctx := context.WithValue(r.Context(), "userEmail", userClaims.Email)
		ctx = context.WithValue(ctx, "username", userClaims.Username)
		ctx = context.WithValue(ctx, "userID", userClaims.UserID)

		// Call next handler with updated context
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
