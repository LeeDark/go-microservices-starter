package server

import (
	"context"

	"github.com/LeeDark/go-microservices-starter/currency/data"
	currencypb "github.com/LeeDark/go-microservices-starter/currency/protos/currency"
	"github.com/hashicorp/go-hclog"
)

type Currency struct {
	log   hclog.Logger
	rates *data.ExchangeRates
}

func NewCurrency(l hclog.Logger, rates *data.ExchangeRates) *Currency {
	return &Currency{log: l, rates: rates}
}

func (c *Currency) GetRate(ctx context.Context, req *currencypb.RateRequest) (*currencypb.RateResponse, error) {
	c.log.Info("GetRate called", "base", req.GetBase(), "destination", req.GetDestination())

	// For demo purposes, return a static rate
	// rate := float32(1.23)
	rate, err := c.rates.GetRate(req.GetBase().String(), req.GetDestination().String())
	if err != nil {
		c.log.Error("Error getting rate", "error", err)
		return nil, err
	}

	return &currencypb.RateResponse{
		Rate: rate,
	}, nil
}
