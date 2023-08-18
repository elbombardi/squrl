package handlers

import (
	"github.com/elbombardi/squrl/src/api_service/api/operations/general"
	"github.com/go-openapi/runtime/middleware"
)

func (handlers *Handlers) HandleHealthcheck(healthcheck general.HealthcheckParams) middleware.Responder {
	return general.NewHealthcheckOK().WithPayload("OK")
}
