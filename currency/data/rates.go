package data

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"strconv"

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

	return nil
}

type Cubes struct {
	CubeData []Cube `xml:"Cube>Cube>Cube"`
}

type Cube struct {
	Currency string `xml:"currency,attr"`
	Rate     string `xml:"rate,attr"`
}
