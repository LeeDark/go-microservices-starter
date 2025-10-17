package handlers

import (
	"net/http"

	"github.com/LeeDark/go-microservices-starter/data"
)

// swagger:route GET /products products listProducts
// Returns a list of products
//
// This will show all available products in the system
//
// Responses:
//   200: productsResponse

// GetProducts returns the products from the data store
func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	lp := data.GetProducts()

	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
}
