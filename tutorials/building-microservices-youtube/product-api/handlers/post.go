package handlers

import (
	"net/http"

	"github.com/LeeDark/go-microservices-starter/tutorials/building-microservices-youtube/product-api/data"
)

// swagger:route POST /products products createProduct
// Create a new product
//
// This will add a new product to the data store
//
// Responses:
//   200: productResponse
//   422: errorValidation
//   501: errorResponse

// Create handles POST requests to add new products
func (p *Products) Create(rw http.ResponseWriter, r *http.Request) {
	// fetch the product from the context
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	p.l.Debug("Inserting product: %#v\n", prod)
	p.productsDB.AddProduct(prod)
}
