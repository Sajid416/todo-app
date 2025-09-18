package product

import (
	"net/http"
	"strconv"
	"strings"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idr := strings.TrimPrefix(r.URL.Path, "/product/")
	id, err := strconv.Atoi(idr)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := h.ProductDB.Exec("DELETE FROM products WHERE id=$1", id)
	if err != nil {
		http.Error(w, "failed to delete task: "+err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "failed to check rows affected: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.Error(w, "Task Not Found", http.StatusNotFound)
	}

	WriteJson(w, http.StatusOK, "Task Deleted Successfully")
}
