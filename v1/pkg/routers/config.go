package routers

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	AuthApiKey string `envconfig:"AUTHORIZATION_API_KEY"`
}

// ConfigFromEnv ...
func ConfigFromEnv(prefix string) (*Config, error) {
	conf := Config{}
	e := envconfig.Process(prefix, &conf)
	return &conf, e
}
