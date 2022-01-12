package coinbasepro

import (
	"fmt"
)

type Currency struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	MinSize      string `json:"min_size"`
	Status       string `json:"status"`
	Message      string `json:"message"`
	MaxPrecision string `json:"max_precision"`
}

func (c *Client) GetCurrencies() ([]Currency, error) {
	var currencies []Currency

	url := fmt.Sprintf("/currencies")
	_, err := c.Request("GET", url, nil, &currencies)
	return currencies, err
}
