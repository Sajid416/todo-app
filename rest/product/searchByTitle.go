package product

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/Sajid416/todo-app/model"
)

func (h *Handler) SearchTaskByTitle(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")

	var product []model.Product

	err := h.middlewares.DB.Select(&product, `select * from products where title ILIKE $1`, "%"+title+"%")
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Product Not Found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	WriteJson(w, http.StatusFound, product)
}
