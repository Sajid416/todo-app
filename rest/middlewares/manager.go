package middlewares

import "net/http"

type Middleware func(http.Handler) http.Handler

type Manager struct {
	middlewares []Middleware
}

func NewManager() *Manager {
	return &Manager{}
}

func (m *Manager) Use(mws ...Middleware) {
	m.middlewares = append(m.middlewares, mws...)
}

func (m *Manager) WrapMux(mux *http.ServeMux) http.Handler {
	var handler http.Handler = mux
	for i := len(m.middlewares) - 1; i >= 0; i-- {
		handler = m.middlewares[i](handler)
	}
	return handler
}
func (m *Manager) WrapHandler(handler http.Handler, mws ...Middleware) http.Handler {
	// Apply middlewares in reverse order so the first added runs last
	for i := len(mws) - 1; i >= 0; i-- {
		handler = mws[i](handler)
	}
	return handler
}