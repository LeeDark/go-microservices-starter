package handlers

import (
	"net/http"

	"github.com/LeeDark/go-microservices-starter/data"
)

// swagger:route POST /products products addProduct
// Adds a new product to the system
//
// This will add a new product to the data store
//
// Responses:
//   200: productResponse
//   422: errorValidation
//   501: errorResponse

// AddProduct adds a new product to the data store
func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	prod := r.Context().Value(KeyProduct{}).(*data.Product)

	p.l.Printf("[DEBUG] Inserting product: %#v\n", prod)

	data.AddProduct(prod)
}
