package handlers

import (
	"github.com/elbombardi/squrl/src/api_service/api/models"
	"github.com/elbombardi/squrl/src/api_service/api/operations/links"
	"github.com/elbombardi/squrl/src/api_service/core"
	"github.com/go-openapi/runtime/middleware"
)

// Handler for the POST /urls endpoint
func (handlers *Handlers) HandleCreateLink(params links.CreateLinkParams, principal any) middleware.Responder {

	if params.Body == nil {
		return links.NewCreateLinkBadRequest().WithPayload(&models.Error{
			Message: "Request body is required",
		})
	}

	shortUrl, err := handlers.LinksManager.Shorten(params.Body.LongURL, principal.(*core.User))

	if err != nil {
		coreError, ok := err.(*core.CoreError)
		switch {
		case ok && coreError.Code == core.ERR_BAD_PARAMS:
			return links.NewCreateLinkBadRequest().WithPayload(&models.Error{
				Message: coreError.Message,
			})
		case ok && coreError.Code == core.ERR_UNAUTHORIZED:
			return links.NewCreateLinkUnauthorized().WithPayload(&models.Error{
				Message: "Unauthorized access"})
		case ok && coreError.Code == core.ERR_ACCOUNT_NOT_FOUND:
			return links.NewCreateLinkUnauthorized().WithPayload(&models.Error{
				Message: "Account not found for this username: " + principal.(*core.User).Username})
		case ok && coreError.Code == core.ERR_ACCOUNT_DISABLED:
			return links.NewCreateLinkUnauthorized().WithPayload(&models.Error{
				Message: "Account disabled"})
		default:
			return links.NewCreateLinkInternalServerError().WithPayload(&models.Error{
				Message: "Internal server error",
			})
		}
	}

	return links.NewCreateLinkOK().WithPayload(&models.URLCreated{
		ShortURL:    shortUrl.ShortUrl.String(),
		ShortURLKey: shortUrl.ShortUrlKey,
	})
}
