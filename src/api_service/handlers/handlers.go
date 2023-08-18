package handlers

import (
	"github.com/elbombardi/squrl/src/api_service/api/models"
	"github.com/elbombardi/squrl/src/api_service/api/operations"
	"github.com/elbombardi/squrl/src/api_service/api/operations/accounts"
	"github.com/elbombardi/squrl/src/api_service/api/operations/general"
	"github.com/elbombardi/squrl/src/api_service/api/operations/urls"
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
	api.GeneralHealthcheckHandler = general.HealthcheckHandlerFunc(handlers.HandleHealthcheck)
	api.GeneralLoginHandler = general.LoginHandlerFunc(handlers.HandleLogin)

	api.AccountsCreateAccountHandler = accounts.CreateAccountHandlerFunc(handlers.HandleCreateAccount)
	api.AccountsUpdateAccountHandler = accounts.UpdateAccountHandlerFunc(handlers.HandleUpdateAccount)

	api.UrlsCreateURLHandler = urls.CreateURLHandlerFunc(handlers.HandleCreateURL)
	api.UrlsUpdateURLHandler = urls.UpdateURLHandlerFunc(handlers.HandleUpdateShortURL)
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

func getError(err error) *models.Error {
	return &models.Error{Error: err.Error()}
}
