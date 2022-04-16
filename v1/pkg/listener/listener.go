package listener

import (
	"context"
	"fmt"
	"log"
	"time"

	apdbabstract "github.com/vallinplasencia/boletia/v1/pkg/external-services/db/abstract"
	apdbclient "github.com/vallinplasencia/boletia/v1/pkg/external-services/db/postgre"
	apexchangeabstract "github.com/vallinplasencia/boletia/v1/pkg/external-services/exchange/abstract"
	apexchangeclient "github.com/vallinplasencia/boletia/v1/pkg/external-services/exchange/client"
)

// Server ...
type Listener struct {
	// db service
	db        apdbabstract.DB
	exchanges apexchangeabstract.Exchanges
}

func New() *Listener {
	projectName := "Boletia"

	// db
	dbConf, e := apdbclient.ConfigFromEnv(projectName)
	we("unable read db environment", e)
	db, e := apdbclient.New(dbConf)
	we("unable configurate db client", e)

	// exchange rates
	exConf, e := apexchangeclient.ConfigFromEnv(projectName)
	we("unable read exchange rates environment", e)
	ex, e := apexchangeclient.New(exConf)
	we("unable configurate exchange rates client", e)

	return &Listener{
		db:        db,
		exchanges: ex,
	}
}

// Run ...
func (s *Listener) Run() {
	for {
		d, code, e := s.exchanges.GetExchangesRates()
		if e != nil {
			fmt.Println("error echange: ", e.Error())
		}
		fmt.Println("Code requesto to echange rates: ", code)
		if d.Meta != nil {
			fmt.Println("LastUpdate: ", d.Meta.LastUpdateAt)
		}
		for k, v := range d.Data {
			fmt.Printf("%s: %+v", k, v)
		}
		fmt.Println("---------------------------------------------------------------------------")
		time.Sleep(10 * time.Second)
	}

}

// Run ...
func (s *Listener) Shutdown(ctx context.Context) {

}

// log error
func we(prefix string, e error) {
	if e != nil && len(prefix) > 0 {
		log.Fatalf("%s, %s", prefix, e.Error())
	} else if e != nil {
		log.Fatal(e)
	}
}
