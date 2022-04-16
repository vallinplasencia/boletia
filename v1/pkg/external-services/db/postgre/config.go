package mysql

import (
	"errors"
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type config struct {
	databaseUrl string
}

// ConfigFromEnv ...
func ConfigFromEnv(prefix string) (*config, error) {
	temp := struct {
		User     string `envconfig:"POSTGRE_DB_USER"`
		Password string `envconfig:"POSTGRE_DB_PASSWORD"`
		// Address incluye el puerto
		Address string `envconfig:"POSTGRE_DB_ADDRESS"`
		DBName  string `envconfig:"POSTGRE_DB_DBNAME"`
	}{}
	e := envconfig.Process(prefix, &temp)
	if len(temp.User) == 0 {
		return nil, errors.New("user is empty")
	}
	if len(temp.Password) == 0 {
		return nil, errors.New("password is empty")
	}
	if len(temp.Address) == 0 {
		return nil, errors.New("address is empty")
	}
	if len(temp.DBName) == 0 {
		return nil, errors.New("dbname is empty")
	}
	return &config{
		databaseUrl: fmt.Sprintf("postgres://%s:%s@%s/%s", temp.User, temp.Password, temp.Address, temp.DBName),
	}, e
}
