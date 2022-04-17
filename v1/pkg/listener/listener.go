package listener

import (
	"context"
	"log"
	"time"

	apdbabstract "github.com/vallinplasencia/boletia/v1/pkg/external-services/db/abstract"
	apdbclient "github.com/vallinplasencia/boletia/v1/pkg/external-services/db/postgres"
	apexchangeabstract "github.com/vallinplasencia/boletia/v1/pkg/external-services/exchange/abstract"
	apexchangeclient "github.com/vallinplasencia/boletia/v1/pkg/external-services/exchange/client"
	apmodels "github.com/vallinplasencia/boletia/v1/pkg/models"
)

// Server ...
type Listener struct {
	// db service
	db        apdbabstract.DB
	exchanges apexchangeabstract.Exchanges

	// delayBetweenCicles minutos q se demora la llamada a la api de echange
	delayBetweenCicles int
}

func New() *Listener {
	projectName := "Boletia"

	// listener
	listenerConf, e := ConfigFromEnv(projectName)
	we("unable read listener environment", e)

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
		db:                 db,
		exchanges:          ex,
		delayBetweenCicles: listenerConf.DelayBetweenCicles,
	}
}

// Run ...
func (s *Listener) Run() {

	go func() {
		numCicle := 1
		for {
			log.Printf("-------------------------------------Start cicle #: %d-------------------------------------", numCicle)

			now := time.Now().UTC()
			d, _, e := s.exchanges.GetExchangesRates()
			duration := time.Now().UTC().Sub(now)
			status := apmodels.StatusOK
			if e != nil {
				log.Println("get echanges, error: ", e.Error())
				status = apmodels.FailedOK
			} else {
				data := []*apmodels.Currency{}

				for _, v := range d.Data {
					data = append(data, &apmodels.Currency{
						Code:         v.Code,
						Value:        v.Value,
						LastUpdateAt: d.LastUpdateAt,
						CreatedAt:    now.Unix(),
						UpdatedAt:    0,
					})
				}
				e = s.db.Currencies().Add(data)
				if e != nil {
					log.Println("save many currencies, error: ", e.Error())
				} else {
					log.Printf("Save %d currencies", len(data))
				}
			}
			e = s.db.Requests().Add(&apmodels.Request{
				Date:      now,
				Duration:  duration.String(),
				Status:    status,
				CreatedAt: now.Unix(),
				UpdatedAt: 0,
			})
			if e != nil {
				log.Println("save many currencies, error: ", e.Error())
			}
			log.Printf("-------------------------------------End cicle #: %d. Wait %d minutes for a new cicle-------------------------------------", numCicle, s.delayBetweenCicles)
			numCicle++
			<-time.After(time.Duration(s.delayBetweenCicles) * time.Minute)
		}
	}()
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
