package v1

import (
	"errors"

	aputil "github.com/vallinplasencia/boletia/v1/pkg/util"
)

// === errors === //

var (
	// errorRespIsEmpty external service response is empty
	errorRespIsEmpty error = errors.New("response is nil")
)

// === concurrency === //
const (
	concFindY aputil.ConcName = "find-y"
)

// === headers === //

const (
	// headerXTotalItems header para el total de items en un listado
	headerXTotalItems string = "X-Total-Items"
)
