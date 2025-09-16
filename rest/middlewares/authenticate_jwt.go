package middlewares

import (
	"net/http"
	"strings"

	"github.com/Sajid416/todo-app/rest/util"
)

func (m *Manager) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			util.SendData(w, "Missing authorization header", http.StatusUnauthorized)
			return
		}
		tokenStr:=strings.TrimPrefix(tokenString,"Bearer ")
		err := util.VerifyToken(tokenStr)		
		
		
		if err != nil {
			util.SendData(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
