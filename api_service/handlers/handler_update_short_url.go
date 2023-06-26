package handlers

import (
	"github.com/elbombardi/squrl/api_service/api/operations"
	"github.com/go-openapi/runtime/middleware"
)

func (*Handlers) UpdateShortURLHandler(params operations.PutShortURLParams) middleware.Responder {
	return middleware.NotImplemented("operation operations.PutShortURL has not yet been implemented")
}

func internalErrorInUpdateShortURL(err error) middleware.Responder {
	return operations.NewPutShortURLInternalServerError().WithPayload(&operations.PutShortURLInternalServerErrorBody{
		Error: err.Error()})
}
