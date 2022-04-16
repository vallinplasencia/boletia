package abstract

// Exchanges ...
type Exchanges interface {
	// GetExchangesRates return echange rates with base USD
	GetExchangesRates() (*ExchangeRates, int, error)
}

// Exchangerates ...
type ExchangeRates struct {
	Meta *struct {
		LastUpdateAt string `json:"last_updated_at"`
	} `json:"meta"`
	Data map[string]*Currency `json:"data"`
}

// Currency ...
type Currency struct {
	Code  string  `json:"code"`
	Value float64 `json:"value"`
}
