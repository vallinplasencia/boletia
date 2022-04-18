package abstract

// Exchanges ...
type Exchanges interface {
	// GetExchangesRates return echange rates with base USD
	GetExchangesRates() (*ExchangeRates, int, error)
}

// Exchangerates ...
type ExchangeRates struct {
	// LastUpdateAt fecha formato unix
	LastUpdateAt int64
	Data         map[string]*Currency
}

// Currency ...
type Currency struct {
	Code  string
	Value float64
}
