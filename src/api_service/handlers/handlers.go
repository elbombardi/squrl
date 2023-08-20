package handlers

import (
	"github.com/elbombardi/squrl/src/api_service/api/operations"
	"github.com/elbombardi/squrl/src/api_service/api/operations/accounts"
	"github.com/elbombardi/squrl/src/api_service/api/operations/general"
	"github.com/elbombardi/squrl/src/api_service/api/operations/links"
	"github.com/elbombardi/squrl/src/api_service/core"
	"github.com/elbombardi/squrl/src/api_service/util"
)

type Handlers struct {
	core.AccountsManager
	core.Authenticator
	core.LinksManager
	*util.Config
}

func (handlers *Handlers) InstallHandlers(api *operations.AdminAPI) {
	api.GeneralHealthcheckHandler = general.HealthcheckHandlerFunc(handlers.HandleHealthcheck)
	api.GeneralLoginHandler = general.LoginHandlerFunc(handlers.HandleLogin)
	api.AccountsCreateAccountHandler = accounts.CreateAccountHandlerFunc(handlers.HandleCreateAccount)
	api.AccountsUpdateAccountHandler = accounts.UpdateAccountHandlerFunc(handlers.HandleUpdateAccount)
	api.LinksCreateLinkHandler = links.CreateLinkHandlerFunc(handlers.HandleCreateLink)
	api.LinksUpdateLinkHandler = links.UpdateLinkHandlerFunc(handlers.HandleUpdateLink)
	api.BearerAuth = func(s string) (any, error) {
		user, _ := handlers.Authenticator.Validate(s)
		return user, nil
	}
}

func encodeStatus(status string) core.Optional[bool] {
	if status == "" {
		return core.Optional[bool]{
			IsSet: false,
		}
	}
	return core.Optional[bool]{
		IsSet: true,
		Value: status == "active",
	}
}

func decodeStatus(enabled bool) string {
	if enabled {
		return "active"
	}
	return "inactive"
}
