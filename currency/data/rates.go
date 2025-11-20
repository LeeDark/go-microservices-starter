package data

import (
	"encoding/xml"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/go-hclog"
)

type ExchangeRates struct {
	log   hclog.Logger
	rates map[string]float64
}

func NewRates(l hclog.Logger) (*ExchangeRates, error) {
	er := &ExchangeRates{
		log:   l,
		rates: map[string]float64{},
	}

	err := er.getRates()

	return er, err
}

func (e *ExchangeRates) GetRate(base, dest string) (float64, error) {
	rate, ok := e.rates[base]
	if !ok {
		return 0, fmt.Errorf("currency %s not found", base)
	}

	destRate, ok := e.rates[dest]
	if !ok {
		return 0, fmt.Errorf("currency %s not found", dest)
	}

	rate = destRate / rate
	if rate == 0 {
		return 0, fmt.Errorf("invalid rate for %s to %s", base, dest)
	}

	return rate, nil
}

// MonitorRates checks the rates in the ECB API every interval and sends a message to the
// returned channel when there are changes
//
// Note: the ECB API only returns data once a day, this function only simulates the changes
// in rates for demonstration purposes
func (e *ExchangeRates) MonitorRates(interval time.Duration) chan struct{} {
	ret := make(chan struct{})

	go func() {
		ticker := time.NewTicker(interval)
		for range ticker.C {
			// just add a random difference to the rate and return it
			// this simulates the fluctuations in currency rates
			for k, v := range e.rates {
				// change can be 10% of original value
				change := (rand.Float64() / 10)
				// is this a postive or negative change
				direction := rand.Intn(2)

				if direction == 0 {
					// new value with be min 90% of old
					change = 1 - change
				} else {
					// new value will be 110% of old
					change = 1 + change
				}

				// modify the rate
				e.rates[k] = v * change
			}

			// notify updates, this will block unless there is a listener on the other end
			ret <- struct{}{}
		}
	}()

	return ret
}

func (e *ExchangeRates) getRates() error {
	resp, err := http.DefaultClient.Get("https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml")
	if err != nil {
		e.log.Error("Unable to get exchange rates", "error", err)
		return err
	}

	if resp.StatusCode != http.StatusOK {
		e.log.Error("Unable to get exchange rates", "status", resp.StatusCode)
		return fmt.Errorf("expected error code 200 got %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	md := &Cubes{}
	err = xml.NewDecoder(resp.Body).Decode(md)
	if err != nil {
		e.log.Error("Unable to decode exchange rates", "error", err)
		return err
	}

	for _, cube := range md.CubeData {
		rate, err := strconv.ParseFloat(cube.Rate, 64)
		if err != nil {
			e.log.Error("Unable to parse rate", "currency", cube.Currency, "rate", cube.Rate, "error", err)
			continue
		}
		e.rates[cube.Currency] = rate
	}

	e.rates["EUR"] = 1.0

	return nil
}

type Cubes struct {
	CubeData []Cube `xml:"Cube>Cube>Cube"`
}

type Cube struct {
	Currency string `xml:"currency,attr"`
	Rate     string `xml:"rate,attr"`
}
