package product

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/Sajid416/todo-app/model"
)

func WriteJson(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)

}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	param := strings.TrimPrefix(r.URL.Path, "/product/")
	id, err := strconv.Atoi(param)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var product model.Product
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	query := `UPDATE products 
          SET title=$1, description=$2, imgUrl=$3 
          WHERE id=$4 
          RETURNING id, title, description, imgUrl`

	err = h.TodoDB.Get(&product, query, product.Title, product.Description, product.ImgUrl, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to update product: "+err.Error(), http.StatusInternalServerError)
		return
	}

	WriteJson(w, http.StatusOK, product)
}
