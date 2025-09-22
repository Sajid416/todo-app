package product

import (
	"net/http"

	"github.com/Sajid416/todo-app/model"
)

func (h *Handler) GetAllProduct(w http.ResponseWriter, r *http.Request) {
	var products []model.Product

	query := `SELECT id, title, description, img_url FROM products ORDER BY id`
	err := h.middlewares.DB.Select(&products, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(products) == 0 {
		http.Error(w, "No products found", http.StatusNotFound)
		return
	}

	WriteJson(w, http.StatusOK, products)

}
