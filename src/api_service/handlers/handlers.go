package handlers

import (
	"github.com/elbombardi/squrl/src/api_service/api/operations"
	"github.com/elbombardi/squrl/src/api_service/util"
	"github.com/elbombardi/squrl/src/db"
)

type Handlers struct {
	db.AccountRepository
	db.URLRepository
	db.ClickRepository
	*util.Config
}

func (handlers *Handlers) InstallHandlers(api *operations.AdminAPI) {
	api.PostAccountHandler = operations.PostAccountHandlerFunc(handlers.HandleCreateAccount)
	api.PostShortURLHandler = operations.PostShortURLHandlerFunc(handlers.HandleCreateURL)
	api.PutAccountHandler = operations.PutAccountHandlerFunc(handlers.HandleUpdateAccount)
	api.PutShortURLHandler = operations.PutShortURLHandlerFunc(handlers.HandleUpdateShortURL)
}

func encodeStatus(status string) bool {
	return status == "active"
}

func decodeStatus(enabled bool) string {
	if enabled {
		return "active"
	}
	return "inactive"
}
