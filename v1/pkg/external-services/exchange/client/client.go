package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	apechangeabstract "github.com/vallinplasencia/boletia/v1/pkg/external-services/exchange/abstract"
)

const (
	headerApiKey string = "apikey"
)

type exchangeClient struct {
	// urlBase del servicio
	urlBase string
	// apiKey api-key de authorization de currencyapi
	apiKey string

	client *http.Client
}

// New ...
func New(c *config) (apechangeabstract.Exchanges, error) {
	return &exchangeClient{
		urlBase: c.URLBase,
		apiKey:  c.ApiKey,
		client: &http.Client{
			Timeout: time.Duration(c.RequestHttpTimeout) * time.Second,
		},
	}, nil
}

// Analyze ...
func (c *exchangeClient) GetExchangesRates() (*apechangeabstract.ExchangeRates, int, error) {
	fullUrl := fmt.Sprintf("%s/latest", c.urlBase)
	req, e := http.NewRequest(
		http.MethodGet,
		fullUrl,
		nil,
	)
	if e != nil {
		return nil, http.StatusInternalServerError, e
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set(headerApiKey, c.apiKey) // authorization

	resp, e := c.client.Do(req)
	if e != nil {
		return nil, http.StatusInternalServerError, e
	}
	defer resp.Body.Close()
	code := resp.StatusCode
	respBody, e := ioutil.ReadAll(resp.Body)

	if code > http.StatusAccepted {
		body := ""
		if e != nil {
			body = fmt.Sprintf("read body error: %s", e)
		} else {
			body = string(respBody)
		}
		return nil, code, fmt.Errorf("problem to request a: %s, code: %d, body: %s", fullUrl, resp.StatusCode, body)
	}
	if e != nil {
		return nil, http.StatusInternalServerError, e
	}
	// RespAnalysis ...
	bResp := exchangeRates{}

	e = json.Unmarshal(respBody, &bResp)
	if e != nil {
		return nil, http.StatusInternalServerError, e
	}
	data := map[string]*apechangeabstract.Currency{}
	for k, v := range bResp.Data {
		data[k] = &apechangeabstract.Currency{
			Code:  v.Code,
			Value: v.Value,
		}
	}
	date, e := time.Parse("2006-01-02T15:04:05Z", bResp.Meta.LastUpdateAt)
	return &apechangeabstract.ExchangeRates{
		LastUpdateAt: date.UTC().Unix(),
		Data:         data,
	}, code, e
}

// exchangerates ...
type exchangeRates struct {
	Meta *struct {
		LastUpdateAt string `json:"last_updated_at"`
	} `json:"meta"`
	Data map[string]*currency `json:"data"`
}

// currency ...
type currency struct {
	Code  string  `json:"code"`
	Value float64 `json:"value"`
}
