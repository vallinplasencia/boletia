package v1

import (
	apdbabstract "github.com/vallinplasencia/boletia/v1/pkg/external-services/db/abstract"
	aputil "github.com/vallinplasencia/boletia/v1/pkg/util"
)

// Base ...
type base struct {
	env aputil.EnvType

	// === external services === //

	// db service
	db apdbabstract.DB
}
