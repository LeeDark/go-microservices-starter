package main

import (
	"testing"

	"github.com/LeeDark/go-microservices-starter/sdk/client"
	"github.com/LeeDark/go-microservices-starter/sdk/client/products"
)

func TestOurClient(t *testing.T) {
	c := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host:     "localhost:9090",
		BasePath: "/",
		Schemes:  []string{"http"},
	})

	// Alternatively, you can use the default client
	// c := client.Default
	// but ensure that the server is running at the expected host and port.

	params := products.NewListProductsParams()
	prod, err := c.Products.ListProducts(params)

	if err != nil {
		t.Fatalf("Error fetching products: %v", err)
	}

	if len(prod.Payload) == 0 {
		t.Fatalf("Expected at least one product, got none")
	}

	t.Logf("Fetched %d products successfully", len(prod.Payload))

	t.Logf("First product: %+v", prod.Payload[0])
}
