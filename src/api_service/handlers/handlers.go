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

func (h *Handlers) InstallHandlers(api *operations.AdminAPI) {
	api.GeneralHealthcheckHandler = general.HealthcheckHandlerFunc(h.HandleHealthcheck)
	api.GeneralLoginHandler = general.LoginHandlerFunc(h.HandleLogin)
	api.AccountsCreateAccountHandler = accounts.CreateAccountHandlerFunc(h.HandleCreateAccount)
	api.AccountsUpdateAccountHandler = accounts.UpdateAccountHandlerFunc(h.HandleUpdateAccount)
	api.LinksCreateLinkHandler = links.CreateLinkHandlerFunc(h.HandleCreateLink)
	api.LinksUpdateLinkHandler = links.UpdateLinkHandlerFunc(h.HandleUpdateLink)
	api.BearerAuth = func(s string) (any, error) {
		user, _ := h.Authenticator.Validate(s)
		return user, nil
	}
}

func encodeStatus(s string) core.Optional[bool] {
	if s == "" {
		return core.Optional[bool]{
			IsSet: false,
		}
	}
	return core.Optional[bool]{
		IsSet: true,
		Value: s == "active",
	}
}

func decodeStatus(enabled bool) string {
	if enabled {
		return "active"
	}
	return "inactive"
}
