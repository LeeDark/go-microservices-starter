package server

import (
	"context"
	"io"
	"time"

	"github.com/LeeDark/go-microservices-starter/tutorials/building-microservices-youtube/currency/data"
	currencypb "github.com/LeeDark/go-microservices-starter/tutorials/building-microservices-youtube/currency/protos/currency"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Currency struct {
	log           hclog.Logger
	rates         *data.ExchangeRates
	subscriptions map[currencypb.Currency_SubscribeRatesServer][]*currencypb.RateRequest
}

func NewCurrency(l hclog.Logger, rates *data.ExchangeRates) *Currency {
	c := &Currency{
		log:           l,
		rates:         rates,
		subscriptions: make(map[currencypb.Currency_SubscribeRatesServer][]*currencypb.RateRequest),
	}

	go c.handleUpdates()

	return c
}

func (c *Currency) handleUpdates() {
	ru := c.rates.MonitorRates(5 * time.Second)
	for range ru {
		c.log.Info("Got updated rates")

		// loop over subscrived clients
		for k, v := range c.subscriptions {

			// loop over subscribed rates
			for _, rr := range v {
				rate, err := c.rates.GetRate(rr.GetBase().String(), rr.GetDestination().String())
				if err != nil {
					c.log.Error("Unable to get updated rate", "base", rr.GetBase().String(), "destination", rr.GetDestination().String(), "error", err)
					continue
				}

				err = k.Send(&currencypb.RateResponse{Base: rr.Base, Destination: rr.Destination, Rate: rate})
				if err != nil {
					c.log.Error("Unable to send updated rate to client", "error", err)
				}
			}
		}
	}
}

func (c *Currency) GetRate(ctx context.Context, req *currencypb.RateRequest) (*currencypb.RateResponse, error) {
	c.log.Info("GetRate called", "base", req.GetBase(), "destination", req.GetDestination())

	if req.Base == req.Destination {
		// err := status.Errorf(
		// 	codes.InvalidArgument,
		// 	"Base currency %s can not be the same as the destination currency %s",
		// 	req.Base.String(), req.Destination.String(),
		// )

		err := status.Newf(
			codes.InvalidArgument,
			"Base currency %s can not be the same as the destination currency %s",
			req.Base.String(), req.Destination.String(),
		)

		err, wde := err.WithDetails(req)
		if wde != nil {
			return nil, wde
		}

		// return nil, fmt.Errorf("base can not be the same as the destination")
		return nil, err.Err()
	}

	// For demo purposes, return a static rate
	// rate := float32(1.23)
	rate, err := c.rates.GetRate(req.GetBase().String(), req.GetDestination().String())
	if err != nil {
		c.log.Error("Error getting rate", "error", err)
		return nil, err
	}

	return &currencypb.RateResponse{
		Base:        req.GetBase(),
		Destination: req.GetDestination(),
		Rate:        rate,
	}, nil
}

func (c *Currency) SubscribeRates(stream currencypb.Currency_SubscribeRatesServer) error {
	for {
		// Recv is a blocking method which returns on client data
		rr, err := stream.Recv()
		// io.EOF signals that the client has cloased the connection
		if err == io.EOF {
			c.log.Info("Client closed the stream")
			break
		}

		// any other error means the transport between the server and client is unavaliable
		if err != nil {
			c.log.Error("Unable to read from client", "error", err)
			return err
		}

		c.log.Info("SubscribeRates called", "base", rr.GetBase(), "destination", rr.GetDestination())

		rrs, ok := c.subscriptions[stream]
		if !ok {
			rrs = []*currencypb.RateRequest{}
		}

		rrs = append(rrs, rr)
		c.subscriptions[stream] = rrs
	}

	return nil
}
