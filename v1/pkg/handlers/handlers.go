package v1

import (
	"github.com/kelseyhightower/envconfig"

	apdbabstract "github.com/vallinplasencia/boletia/v1/pkg/external-services/db/abstract"
	aputil "github.com/vallinplasencia/boletia/v1/pkg/util"
)

// Handlers ...
type Handlers struct {
	Currencies *CurrenciesHandler
}

func New(c *config, db apdbabstract.DB) *Handlers {
	base := &base{
		db:  db,
		env: c.Env,
	}
	return &Handlers{
		Currencies: &CurrenciesHandler{
			base: base,
		},
	}
}

// ======= config ======= //

type config struct {
	Env aputil.EnvType
}

// ConfigFromEnv ...
func ConfigFromEnv(prefix string) (*config, error) {
	temp := struct {
		Env string `envconfig:"APP_MODE"`
	}{}
	e := envconfig.Process(prefix, &temp)

	conf := config{
		Env: aputil.EnvProduction,
	}
	switch temp.Env {
	case string(aputil.EnvProduction):
		conf.Env = aputil.EnvProduction
	case string(aputil.EnvDevelop):
		conf.Env = aputil.EnvDevelop
	}
	return &conf, e
}
