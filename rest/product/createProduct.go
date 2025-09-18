package product

import (
	"encoding/json"
	"net/http"

	"github.com/Sajid416/todo-app/model"
)

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var product model.Product

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	query := `INSERT INTO products (title,description,imgUrl) VALUES ($1, $2,$3) RETURNING id,title,description,imgUrl`
	err := h.ProductDB.Get(&product, query, product.Title, product.Description, product.ImgUrl)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	WriteJson(w, http.StatusCreated, product)

}
