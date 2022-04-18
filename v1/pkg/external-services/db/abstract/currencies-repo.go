package abstract

import (
	apmodels "github.com/vallinplasencia/boletia/v1/pkg/models"
)

// CurrenciesRepo db access
type CurrenciesRepo interface {
	// Add add a new item
	Add(d []*apmodels.Currency) error
	// List list currencies with filters
	// si filter.Finit es cero NO se aplica ese filtro
	// si filter.Fend es cero NO se aplica ese filtro
	// si filter.Currency esta vacio NO se aplica para el filtro
	List(f *FilterCurrencies) ([]*apmodels.Currency, error)
}

// FilterCurrencies filter ..
type FilterCurrencies struct {
	Currency string
	Finit    int64
	Fend     int64
}
