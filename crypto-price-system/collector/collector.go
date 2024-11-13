package collector

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type CryptoPrice struct {
	Currency string  `json:"currency"`
	Price    float64 `json:"price"`
}

func FetchPrices(apiUrl string) ([]CryptoPrice, error) {
	resp, err := http.Get(apiUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var prices []CryptoPrice
	if err := json.Unmarshal(body, &prices); err != nil {
		return nil, err
	}

	return prices, nil
}
