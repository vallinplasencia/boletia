package req

// CurrenciesFilter filtros para list currencies.(Query String)
type CurrenciesFilter struct {
	Finit string `form:"finit"`
	Fend  string `form:"fend"`
}
