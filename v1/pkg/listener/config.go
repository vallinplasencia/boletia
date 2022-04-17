package listener

import (
	"errors"

	"github.com/kelseyhightower/envconfig"
)

type config struct {
	// delayBetweenCicles minutos q se demora entre llamadas
	DelayBetweenCicles int `envconfig:"DELAY_BETWEEN_CICLES"`
}

// ConfigFromEnv ...
func ConfigFromEnv(prefix string) (*config, error) {
	conf := config{}
	e := envconfig.Process(prefix, &conf)
	if conf.DelayBetweenCicles < 0 {
		return nil, errors.New("delay between cicles is invalid")
	}
	return &conf, e
}
