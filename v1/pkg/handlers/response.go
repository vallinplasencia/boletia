package v1

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	aphresp "github.com/vallinplasencia/boletia/v1/pkg/handlers/models/resp"
	aputil "github.com/vallinplasencia/boletia/v1/pkg/util"
)

// response ...
type response struct {
	env aputil.EnvType
	c   *gin.Context
}

// send ...
func (r *response) send(httpCode int, code aphresp.CodeType, data interface{}, headers map[string]string) {
	for k, v := range headers {
		r.c.Header(k, v)
	}
	if httpCode < 400 {
		if data != nil {
			r.c.JSON(httpCode, data)
		} else {
			r.c.JSON(httpCode, nil)
		}
	} else {
		msg := aphresp.GetMsgError(code) // mensaje por defecto
		// solo si esta en modo desarrollo se envia el error con todos sus detalles
		if r.env == aputil.EnvDevelop {
			switch t := data.(type) {
			case string:
				msg = fmt.Sprintf("%s --- %s", msg, t)
			case error:
				msg = fmt.Sprintf("%s --- %s", msg, t.Error())
			default:
				if data != nil {
					msg = fmt.Sprintf("%s --- %+v", msg, data)
				}
			}
			log.Println(msg)
		} else {
			log.Printf("%s. %+v", msg, data)
		}
		r.c.JSON(httpCode, &aphresp.Error{
			Code: code,
			Msg:  msg,
		})
	}
}

// sendOK ...
func (r *response) sendOK(data interface{}, headers map[string]string) {
	r.send(http.StatusOK, aphresp.CodeOK, data, headers)
}

// sendBadRequest ...
func (r *response) sendBadRequest(code aphresp.CodeType, e interface{}, headers map[string]string) {
	r.send(http.StatusBadRequest, code, e, headers)
}

// sendInternalError ...
func (r *response) sendInternalError(code aphresp.CodeType, e interface{}, headers map[string]string) {
	r.send(http.StatusInternalServerError, code, e, headers)
}
