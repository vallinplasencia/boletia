package client

import (
	"errors"
	"strings"

	"github.com/kelseyhightower/envconfig"
)

// config ...
type config struct {
	// URLBase url base para la api de currencyapi
	URLBase            string `envconfig:"CURRENCY_API_URL_BASE"`
	RequestHttpTimeout int    `envconfig:"CURRENCY_REQUEST_TIMEOUT"`
	// ApiKey key de authorization de currencyapi
	ApiKey string `envconfig:"CURRENCY_API_KEY"`
}

// ConfigFromEnv ...
func ConfigFromEnv(prefix string) (*config, error) {
	c := config{}
	e := envconfig.Process(prefix, &c)
	if len(c.ApiKey) == 0 {
		return nil, errors.New("api key is empty")
	}
	c.URLBase = strings.TrimSuffix(c.URLBase, "/")
	if len(c.URLBase) == 0 {
		return nil, errors.New("url base is empty")
	}
	return &c, e
}
