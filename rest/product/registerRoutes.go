package product

import (
	"net/http"

	"github.com/Sajid416/todo-app/rest/middlewares"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager) {

	mux.Handle("GET /product", manager.With(
		http.HandlerFunc(h.GetAllProduct),
	))
	mux.Handle("GET /product/filter", manager.With(
		http.HandlerFunc(h.FilteredProduct),
	))
	mux.Handle("GET /product/{id}", manager.With(
		http.HandlerFunc(h.GetProductById),
	))
	mux.Handle("POST /product", manager.With(
		http.HandlerFunc(h.CreateProduct),
	))
	mux.Handle("PUT /product/{id}", manager.With(
		http.HandlerFunc(h.UpdateProduct),
	))
	mux.Handle("DELETE /product/{id}", manager.With(
		http.HandlerFunc(h.DeleteProduct),
	))

}
