package data

import (
	"testing"

	"github.com/hashicorp/go-hclog"
)

func TestGetRates(t *testing.T) {
	logger := hclog.New(&hclog.LoggerOptions{
		Name:  "test",
		Level: hclog.LevelFromString("DEBUG"),
	})

	tr, err := NewRates(logger)
	if err != nil {
		t.Fatalf("Failed to create ExchangeRates: %v", err)
	}

	// err = tr.getRates()
	// if err != nil {
	// 	t.Fatalf("getRates failed: %v", err)
	// }

	if len(tr.rates) == 0 {
		t.Fatal("No exchange rates were fetched")
	}

	for currency, rate := range tr.rates {
		t.Logf("Currency: %s, Rate: %f", currency, rate)
	}
}
