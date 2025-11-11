package main

import (
	"context"
	"net"
	"os"

	currencypb "github.com/LeeDark/go-microservices-starter/currency/protos/currency"
	"github.com/LeeDark/go-microservices-starter/currency/server"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type currencyServerWrapper struct {
	currencypb.UnimplementedCurrencyServer
	*server.Currency
}

func (w *currencyServerWrapper) GetRate(ctx context.Context, req *currencypb.RateRequest) (*currencypb.RateResponse, error) {
	return w.Currency.GetRate(ctx, req)
}

func main() {
	log := hclog.Default()

	gs := grpc.NewServer()
	cs := server.NewCurrency(log)

	// wrap server.Currency so it satisfies the generated interface that
	// requires embedding UnimplementedCurrencyServer
	currencypb.RegisterCurrencyServer(gs, &currencyServerWrapper{Currency: cs})

	reflection.Register(gs)

	l, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Error("failed to listen", "error", err)
		os.Exit(1)
	}

	log.Info("starting currency gRPC server on :9092")

	gs.Serve(l)
}
