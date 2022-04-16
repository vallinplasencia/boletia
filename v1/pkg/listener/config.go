package listener

import (
	"github.com/kelseyhightower/envconfig"
)

type config struct {
}

// ConfigFromEnv ...
func ConfigFromEnv(prefix string) (*config, error) {
	conf := config{}
	e := envconfig.Process(prefix, &conf)
	return &conf, e
}
