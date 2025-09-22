package product

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/Sajid416/todo-app/model"
)

func (h *Handler) GetProductById(w http.ResponseWriter, r *http.Request) {
	idr := strings.TrimPrefix(r.URL.Path, "/product/")
	id, err := strconv.Atoi(idr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var product model.Product
	err = h.middlewares.DB.Get(&product, `select * from products where id=$1`, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	WriteJson(w, http.StatusOK, product)

}
