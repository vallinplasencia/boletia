package postgres

import (
	"database/sql"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"

	apdbabstract "github.com/vallinplasencia/boletia/v1/pkg/external-services/db/abstract"
)

const (
	schema          string = "public"
	currenciesTable string = "currencies"
	requestsTable   string = "requests"
)

type DB struct {
	currencies apdbabstract.CurrenciesRepo
	requests   apdbabstract.RequestsRepo
}

func New(c *config) (apdbabstract.DB, error) {
	dba, e := sql.Open("pgx", c.databaseUrl)
	if e != nil {
		return nil, e
	}
	dba.SetConnMaxLifetime(time.Minute * 3)
	dba.SetMaxOpenConns(10)
	dba.SetMaxIdleConns(10)

	if e = dba.Ping(); e != nil {
		return nil, e
	}
	return &DB{
		currencies: &currenciesRepo{
			db: dba,
		},
		requests: &requestsRepo{
			db: dba,
		},
	}, nil
}

func (db *DB) Currencies() apdbabstract.CurrenciesRepo {
	return db.currencies
}

func (db *DB) Requests() apdbabstract.RequestsRepo {
	return db.requests
}
