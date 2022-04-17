package postgres

import (
	// database/sql"

	apdbabstract "github.com/vallinplasencia/boletia/v1/pkg/external-services/db/abstract"
)

const currenciesTable string = "currencies"

type DB struct {
	currencies apdbabstract.CurrenciesRepo
}

func New(c *config) (apdbabstract.DB, error) {
	db := &DB{
		currencies: currenciesRepo{},
	}
	return db, nil
}

func (db *DB) Currencies() apdbabstract.CurrenciesRepo {
	return db.currencies
}
