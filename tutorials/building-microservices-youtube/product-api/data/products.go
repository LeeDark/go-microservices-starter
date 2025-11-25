package data

import (
	"context"
	"fmt"

	currencypb "github.com/LeeDark/go-microservices-starter/tutorials/building-microservices-youtube/currency/protos/currency"
	"github.com/hashicorp/go-hclog"
)

// ErrProductNotFound is an error raised when a product can not be found in the database
var ErrProductNotFound = fmt.Errorf("Product not found")

// Product defines the structure for an API product
// swagger:model
type Product struct {
	// the id for the product
	//
	// required: false
	// min: 1
	ID int `json:"id"` // Unique identifier for the product

	// the name for this product
	//
	// required: true
	// max length: 255
	Name string `json:"name" validate:"required"`

	// the description for this product
	//
	// required: false
	// max length: 10000
	Description string `json:"description"`

	// the price for the product
	//
	// required: true
	// min: 0.01
	Price float64 `json:"price" validate:"required,gt=0"`

	// the SKU for the product
	//
	// required: true
	// pattern: [a-z]+-[a-z]+-[a-z]+
	SKU string `json:"sku" validate:"sku"`
}

// Products defines a slice of Product
type Products []*Product

type ProductsDB struct {
	log      hclog.Logger
	currency currencypb.CurrencyClient
	rates    map[string]float64
	client   currencypb.Currency_SubscribeRatesClient
}

func NewProductsDB(l hclog.Logger, c currencypb.CurrencyClient) *ProductsDB {
	db := &ProductsDB{
		log:      l,
		currency: c,
		rates:    make(map[string]float64),
		client:   nil,
	}

	go db.handleUpdates()

	return db
}

func (db *ProductsDB) handleUpdates() {
	sub, err := db.currency.SubscribeRates(context.Background())
	if err != nil {
		db.log.Error("Unable to subscribe for rates", "error", err)
		return
	}

	db.client = sub

	for {
		rr, err := sub.Recv()
		if err != nil {
			db.log.Error("Error receiving message", "error", err)
			continue
		}

		db.log.Info("Received updated rate", "base", rr.GetBase(), "destination", rr.GetDestination(), "rate", rr.GetRate())
		db.rates[rr.GetDestination().String()] = rr.GetRate()
	}
}

// GetProducts returns all products from the database
func (db *ProductsDB) GetProducts(currency string) (Products, error) {
	if currency == "" {
		return productList, nil
	}

	// get exchange rate
	rate, err := db.getRate(currency)
	if err != nil {
		return nil, err
	}

	// productListReturn := make(Products, len(productList))
	// copy(productListReturn, productList)

	productListReturn := Products{}
	for _, prod := range productList {
		np := *prod
		np.Price = np.Price * rate
		productListReturn = append(productListReturn, &np)
	}

	return productListReturn, nil
}

// GetProductByID returns a single product which matches the id from the
// database.
// If a product is not found this function returns a ProductNotFound error
func (db *ProductsDB) GetProductByID(id int, currency string) (*Product, error) {
	i := db.findIndexByProductID(id)
	if id == -1 {
		return nil, ErrProductNotFound
	}

	if currency == "" {
		return productList[i], nil
	}

	rate, err := db.getRate(currency)
	if err != nil {
		db.log.Error("Unable to get rate", "currency", currency, "error", err)
		return nil, err
	}

	np := *productList[i]
	np.Price = np.Price * rate

	return &np, nil
}

// UpdateProduct replaces a product in the database with the given
// item.
// If a product with the given id does not exist in the database
// this function returns a ProductNotFound error
func (db *ProductsDB) UpdateProduct(p Product) error {
	i := db.findIndexByProductID(p.ID)
	if i == -1 {
		return ErrProductNotFound
	}

	// update the product in the DB
	productList[i] = &p

	return nil
}

// AddProduct adds a new product to the database
func (db *ProductsDB) AddProduct(p Product) {
	// get the next id in sequence
	maxID := productList[len(productList)-1].ID
	p.ID = maxID + 1
	productList = append(productList, &p)
}

// DeleteProduct deletes a product from the database
func (db *ProductsDB) DeleteProduct(id int) error {
	i := db.findIndexByProductID(id)
	if i == -1 {
		return ErrProductNotFound
	}

	productList = append(productList[:i], productList[i+1])

	return nil
}

// findIndex finds the index of a product in the database
// returns -1 when no product can be found
func (db *ProductsDB) findIndexByProductID(id int) int {
	for i, p := range productList {
		if p.ID == id {
			return i
		}
	}

	return -1
}

func (db *ProductsDB) getRate(destination string) (float64, error) {
	if r, ok := db.rates[destination]; ok {
		return r, nil
	}

	req := &currencypb.RateRequest{
		Base:        currencypb.Currencies(currencypb.Currencies_value["EUR"]),
		Destination: currencypb.Currencies(currencypb.Currencies_value[destination]),
	}

	// execute gRPC call
	resp, err := db.currency.GetRate(context.Background(), req)
	if err != nil {
		db.log.Error("[ERROR] calling currency service", err)
		return 0, err
	}

	db.rates[destination] = resp.Rate

	// subscribe for updates
	err = db.client.Send(req)
	if err != nil {
		db.log.Error("Unable to send rate request for updates", "error", err)
	}

	return resp.GetRate(), nil
}

// productList is a hard coded list of products for this
// example data source

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
	},
	&Product{
		ID:          2,
		Name:        "Esspresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
	},
}
