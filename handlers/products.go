// Package classification of Product API
//
// # Documentation for Product API
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
// swagger:meta
package handlers

import (
	"log"
)

// KeyProduct is a key used for the product in the context
type KeyProduct struct{}

// Products defines the product handler
type Products struct {
	l *log.Logger
}

// NewProducts creates a new products handler with the given logger
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}
