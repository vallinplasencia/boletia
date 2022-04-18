package abstract

import (
	apmodels "github.com/vallinplasencia/boletia/v1/pkg/models"
)

// RequestsRepo db access
type RequestsRepo interface {
	// Add add a new item
	Add(d *apmodels.Request) error
}
