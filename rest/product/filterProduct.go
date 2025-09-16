package product

import (
	"net/http"

	"github.com/Sajid416/todo-app/model"
)

func (h *Handler) FilteredProduct(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	query := `select * from products where title=$1`
	var products []model.Product
	err := h.TodoDB.Select(&products, query, title)
	if err != nil {
		http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if len(products) == 0 {
		http.Error(w, "Product Not Found", http.StatusNotFound)
		return
	}
	WriteJson(w, http.StatusFound, products)
}
