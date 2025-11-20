package server

import (
	"context"
	"io"
	"time"

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

func (c *Currency) SubscribeRates(stream currencypb.Currency_SubscribeRatesServer) error {

	go func() {
		for {
			rr, err := stream.Recv()
			if err == io.EOF {
				c.log.Info("Client closed the stream")
				break
			}

			if err != nil {
				c.log.Error("Unable to read from client", "error", err)
				break
			}

			c.log.Info("SubscribeRates called", "base", rr.GetBase(), "destination", rr.GetDestination())
		}
	}()

	// For demo purposes, send a rate every 5 seconds
	for {
		err := stream.Send(&currencypb.RateResponse{Rate: 12.1})
		if err != nil {
			c.log.Error("Unable to send to client", "error", err)
			return err
		}

		time.Sleep(5 * time.Second)
	}
}
