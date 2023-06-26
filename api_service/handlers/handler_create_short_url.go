package handlers

import (
	"github.com/elbombardi/squrl/api_service/api/operations"
	"github.com/go-openapi/runtime/middleware"
)

func (*Handlers) CreateShortURLHandler(params operations.PostShortURLParams) middleware.Responder {
	return middleware.NotImplemented("operation operations.PostShortURL has not yet been implemented")
}
