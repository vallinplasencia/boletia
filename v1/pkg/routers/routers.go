package routers

import (
	"github.com/gin-gonic/gin"

	aphand "github.com/vallinplasencia/boletia/v1/pkg/handlers"
	apmiddlewares "github.com/vallinplasencia/boletia/v1/pkg/middlewares"
)

// Router ...
type Router struct {
	eng *gin.Engine
	h   *aphand.Handlers
}

func New(c *Config, eng *gin.Engine, handlers *aphand.Handlers) *Router {
	return &Router{
		eng: eng,
		h:   handlers,
	}
}

// InitRouters set endpoints with yours handlers
func (r *Router) InitRouters() {
	// init v1 routers
	r.initV1Routers()
}

// initV1Routers set endpoints with yours handlers
func (r *Router) initV1Routers() {
	r.eng.Use(apmiddlewares.Cors())
	apiv1 := r.eng.Group("/api/v1")
	{
		apiv1.GET("/currencies/:currency", r.h.Currencies.GetListCurrencies)
	}
}
