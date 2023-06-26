package handlers

import (
	"github.com/elbombardi/squrl/api_service/api/operations"
	db "github.com/elbombardi/squrl/db/sqlc"
)

type Handlers struct {
	CustomersRepository db.CustomersRepository
	ShortURLsRepository db.ShortURLsRepository
	ClicksRepository    db.ClicksRepository
}

func (handlers *Handlers) InstallHandlers(api *operations.ShortURLAPI) {
	api.PostCustomerHandler = operations.PostCustomerHandlerFunc(handlers.CreateCustomerHandler)
	api.PostShortURLHandler = operations.PostShortURLHandlerFunc(handlers.CreateShortURLHandler)
	api.PutCustomerHandler = operations.PutCustomerHandlerFunc(handlers.UpdateCustomerHandler)
	api.PutShortURLHandler = operations.PutShortURLHandlerFunc(handlers.UpdateShortURLHandler)
}
