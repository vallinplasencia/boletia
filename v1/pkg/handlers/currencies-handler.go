package v1

import (
	"github.com/gin-gonic/gin"
)

// CurrenciesHandler incoming request for currencies
type CurrenciesHandler struct {
	*base
}

// GetListCurrencies list currencies with or without filters
func (h *CurrenciesHandler) GetListCurrencies(c *gin.Context) {

}
