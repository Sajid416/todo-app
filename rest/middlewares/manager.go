package middlewares

import "net/http"

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalmiddlewares []Middleware
}

func NewManager() *Manager {
	return &Manager{}
}

func (m *Manager) Use(middlewares ...Middleware) {
	m.globalmiddlewares = append(m.globalmiddlewares, middlewares...)
}

func (mngr *Manager) With(handler http.Handler, middlewares ...Middleware) http.Handler {
    h := handler

    for _, middleware := range middlewares {
        h = middleware(h)
    }

    return h
}

func (mngr *Manager) WrapMux(handler http.Handler) http.Handler {
    h := handler

    for _, middleware := range mngr.globalmiddlewares {
        h = middleware(h)
    }

    return h
}