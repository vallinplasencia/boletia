package postgres

import (
	"database/sql"
	"fmt"
	"strings"

	apdbabstract "github.com/vallinplasencia/boletia/v1/pkg/external-services/db/abstract"
	apmodels "github.com/vallinplasencia/boletia/v1/pkg/models"
)

// currenciesRepo db access
type currenciesRepo struct {
	db *sql.DB
}

// Add add a new item
func (r *currenciesRepo) Add(d []*apmodels.Currency) error {
	values := []string{}
	args := []interface{}{}
	i := 0
	for _, v := range d {
		values = append(values, fmt.Sprintf("($%d,$%d,$%d,$%d,$%d)", i+1, i+2, i+3, i+4, i+5))
		args = append(args, v.Code, v.Value, v.LastUpdateAt, v.CreatedAt, v.UpdatedAt)
		i += 5
	}
	q := fmt.Sprintf("INSERT INTO %s.%s (code,value,last_update_at,created_at,updated_at) VALUES %s", schema, currenciesTable, strings.Join(values, ","))
	_, e := r.db.Exec(q, args...)
	return e
}

// List list currencies with filters
// si filter.Finit es cero NO se aplica ese filtro
// si filter.Fend es cero NO se aplica ese filtro
// si filter.Currency esta vacio NO se aplica para el filtro
func (r *currenciesRepo) List(f *apdbabstract.FilterCurrencies) ([]*apmodels.Currency, error) {
	items := []*apmodels.Currency{}

	params := []interface{}{}
	clauseFilterWhere := ""
	if f != nil {
		conditions := []string{}
		pos := 0
		if len(f.Currency) > 0 {
			pos++
			// conditions = append(conditions, fmt.Sprintf("code=$%d", pos))
			conditions = append(conditions, fmt.Sprintf("code=$%d", pos))
			params = append(params, f.Currency)
		}
		if f.Finit > 0 {
			pos++
			conditions = append(conditions, fmt.Sprintf("last_update_at>=$%d", pos))
			params = append(params, f.Finit)
		}
		if f.Fend > 0 {
			pos++
			conditions = append(conditions, fmt.Sprintf("last_update_at<=$%d", pos))
			params = append(params, f.Fend)
		}
		if len(conditions) > 0 {
			clauseFilterWhere = fmt.Sprintf("%s %s", "WHERE", strings.Join(conditions, " AND "))
		}
	}
	q := fmt.Sprintf("SELECT c.id, c.code, c.value, c.last_update_at, c.created_at ,c.updated_at FROM %s.%s c %s", schema, currenciesTable, clauseFilterWhere)
	result, e := r.db.Query(q, params...)

	if e != nil {
		return items, e
	}
	for result.Next() {
		item := apmodels.Currency{
			ID:           0,
			Code:         "",
			Value:        0,
			LastUpdateAt: 0,
			CreatedAt:    0,
			UpdatedAt:    0,
		}
		e = result.Scan(
			&item.ID,
			&item.Code,
			&item.Value,
			&item.LastUpdateAt,
			&item.CreatedAt,
			&item.UpdatedAt)
		if e != nil {
			return items, e
		}
		items = append(items, &item)
	}
	return items, nil
}
