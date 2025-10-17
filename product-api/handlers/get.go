package handlers

import (
	"net/http"

	"github.com/LeeDark/go-microservices-starter/data"
)

// swagger:route GET /products products listProducts
// Return a list of products from the database
//
// This will show all available products in the system
//
// Responses:
//   200: productsResponse

// ListAll handles GET requests and returns all current products
func (p *Products) ListAll(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("[DEBUG] get all records")

	prods := data.GetProducts()

	err := data.ToJSON(prods, rw)
	if err != nil {
		// we should never be here but log the error just incase
		p.l.Println("[ERROR] serializing product", err)
	}
}
