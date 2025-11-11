package server

import (
	"context"

	currencypb "github.com/LeeDark/go-microservices-starter/currency/protos/currency"
	"github.com/hashicorp/go-hclog"
)

type Currency struct {
	log hclog.Logger
}

func NewCurrency(l hclog.Logger) *Currency {
	return &Currency{log: l}
}

func (c *Currency) GetRate(ctx context.Context, req *currencypb.RateRequest) (*currencypb.RateResponse, error) {
	c.log.Info("GetRate called", "base", req.GetBase(), "destination", req.GetDestination())

	// For demo purposes, return a static rate
	rate := float32(1.23)

	return &currencypb.RateResponse{
		Rate: rate,
	}, nil
}
