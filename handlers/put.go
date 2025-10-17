package handlers

import (
	"net/http"
	"strconv"

	"github.com/LeeDark/go-microservices-starter/data"
	"github.com/gorilla/mux"
)

// swagger:route PUT /products/{id} products updateProduct
// Updates a product in the system
//
// This will update a product in the data store
//
// Responses:
//   204: noContentResponse
//   404: errorResponse
//   422: errorValidation

// UpdateProduct updates a product in the data store
func (p *Products) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert id", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle PUT Product", id)

	prod := r.Context().Value(KeyProduct{}).(*data.Product)

	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}
}
