package v1

import (
	"errors"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	apdbabstract "github.com/vallinplasencia/boletia/v1/pkg/external-services/db/abstract"
	aphreq "github.com/vallinplasencia/boletia/v1/pkg/handlers/models/req"
	aphresp "github.com/vallinplasencia/boletia/v1/pkg/handlers/models/resp"
	apmodels "github.com/vallinplasencia/boletia/v1/pkg/models"
)

// CurrenciesHandler incoming request for currencies
type CurrenciesHandler struct {
	*base
}

// GetListCurrencies list currencies with or without filters
func (h *CurrenciesHandler) GetListCurrencies(c *gin.Context) {
	resp := response{c: c, env: h.env}

	var e error
	f := aphreq.CurrenciesFilter{}
	if e = c.ShouldBindWith(&f, binding.Query); e != nil {
		resp.sendBadRequest(aphresp.CodeInvalidArgument, e, nil)
		return
	}
	pCurrency := strings.ToUpper(c.Param("currency")) // el paramtro "currency" simpre viene en el path de la url
	if pCurrency == "ALL" {
		pCurrency = ""
	}
	dbFilter := &apdbabstract.FilterCurrencies{
		Currency: pCurrency,
		Finit:    0,
		Fend:     0,
	}
	if len(f.Finit) != 0 {
		finit, e := time.Parse("2006-01-02T15:04:05", f.Finit)
		if e != nil {
			resp.sendBadRequest(aphresp.CodeInvalidArgument, e, nil)
			return
		}
		dbFilter.Finit = finit.UTC().Unix()
	}
	if len(f.Fend) != 0 {
		fend, e := time.Parse("2006-01-02T15:04:05", f.Fend)
		if e != nil {
			resp.sendBadRequest(aphresp.CodeInvalidArgument, e, nil)
			return
		}
		dbFilter.Fend = fend.UTC().Unix()
	}
	if len(f.Finit) != 0 && len(f.Fend) != 0 && dbFilter.Finit > dbFilter.Fend {
		resp.sendBadRequest(aphresp.CodeInvalidArgument, errors.New("finit more than fend"), nil)
		return
	}
	items, e := h.db.Currencies().List(dbFilter)
	if e != nil {
		resp.sendInternalError(aphresp.CodeInternalError, e, nil)
		return
	}
	resp.sendOK(h.toRespCurrenciesFromModel(items), nil)
}

// === conv from db === //

func (h *CurrenciesHandler) toRespCurrencyFromModel(d *apmodels.Currency) *aphresp.Currency {
	return &aphresp.Currency{
		ID:           d.ID,
		Code:         d.Code,
		Value:        d.Value,
		LastUpdateAt: d.LastUpdateAt,
		CreatedAt:    d.CreatedAt,
		UpdatedAt:    d.UpdatedAt,
	}
}

func (h *CurrenciesHandler) toRespCurrenciesFromModel(d []*apmodels.Currency) []*aphresp.Currency {
	items := []*aphresp.Currency{}
	for _, v := range d {
		items = append(items, h.toRespCurrencyFromModel(v))
	}
	return items
}
