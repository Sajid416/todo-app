package product

import (
	"net/http"

	"github.com/Sajid416/todo-app/rest/middlewares"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager) {

	mux.Handle("GET /product", manager.WrapHandler(
		http.HandlerFunc(h.GetAllProduct),
	))
	mux.Handle("GET /product/filter", manager.WrapHandler(
		http.HandlerFunc(h.FilteredProduct),
	))
	mux.Handle("GET /product/{id}", manager.WrapHandler(
		http.HandlerFunc(h.GetProductById),
	))
	mux.Handle("POST /product", manager.WrapHandler(
		http.HandlerFunc(h.CreateProduct),
	))
	mux.Handle("PUT /product/{id}", manager.WrapHandler(
		http.HandlerFunc(h.UpdateProduct),
	))
	mux.Handle("DELETE /product/{id}", manager.WrapHandler(
		http.HandlerFunc(h.DeleteProduct),
	))

}
