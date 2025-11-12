package handlers

import (
	"context"
	"net/http"

	currencypb "github.com/LeeDark/go-microservices-starter/currency/protos/currency"
	"github.com/LeeDark/go-microservices-starter/product-api/data"
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
	rw.Header().Add("Content-Type", "application/json")

	prods := data.GetProducts()

	err := data.ToJSON(prods, rw)
	if err != nil {
		// we should never be here but log the error just incase
		p.l.Println("[ERROR] serializing product", err)
	}
}

// swagger:route GET /products/{id} products listSingleProduct
// Return a list of products from the database
// Responses:
//	200: productResponse
//	404: errorResponse

// ListSingle handles GET requests
func (p *Products) ListSingle(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	id := getProductID(r)

	p.l.Println("[DEBUG] get record id", id)

	prod, err := data.GetProductByID(id)

	switch err {
	case nil:

	case data.ErrProductNotFound:
		p.l.Println("[ERROR] fetching product", err)

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	default:
		p.l.Println("[ERROR] fetching product", err)

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	// get exchange rate
	req := &currencypb.RateRequest{
		Base:        currencypb.Currencies(currencypb.Currencies_value["EUR"]),
		Destination: currencypb.Currencies(currencypb.Currencies_value["GBP"]),
	}

	// execute gRPC call
	resp, err := p.cc.GetRate(context.Background(), req)
	if err != nil {
		p.l.Println("[ERROR] calling currency service", err)

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: "error calling currency service"}, rw)
		return
	}

	p.l.Println("[DEBUG] exchange rate response", resp.GetRate())

	// EUR to GBP conversion
	prod.Price = prod.Price * resp.GetRate()

	err = data.ToJSON(prod, rw)
	if err != nil {
		// we should never be here but log the error just incase
		p.l.Println("[ERROR] serializing product", err)
	}
}
