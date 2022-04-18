package postgres

import (
	"database/sql"
	"fmt"

	apmodels "github.com/vallinplasencia/boletia/v1/pkg/models"
)

// requestsRepo db access
type requestsRepo struct {
	db *sql.DB
}

// Add add a new item
func (r *requestsRepo) Add(d *apmodels.Request) error {
	q := fmt.Sprintf("INSERT INTO %s.%s (date,duration,status,created_at,updated_at) VALUES ($1,$2,$3,$4,$5)", schema, requestsTable)
	_, e := r.db.Exec(q, d.Date.Format("2006-01-02T15:04:05"), d.Duration, d.Status, d.CreatedAt, d.UpdatedAt)
	return e
}
