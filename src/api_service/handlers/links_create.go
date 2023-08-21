package handlers

import (
	"github.com/elbombardi/squrl/src/api_service/api/models"
	"github.com/elbombardi/squrl/src/api_service/api/operations/links"
	"github.com/elbombardi/squrl/src/api_service/core"
	"github.com/go-openapi/runtime/middleware"
)

// Handler for the POST /urls endpoint
func (h *Handlers) HandleCreateLink(params links.CreateLinkParams, principal any) middleware.Responder {

	if params.Body == nil {
		return links.NewCreateLinkBadRequest().WithPayload(&models.Error{
			Message: "Request body is required",
		})
	}

	link, err := h.LinksManager.Shorten(params.Body.LongURL, principal.(*core.User))

	if err != nil {
		coreErr, ok := err.(core.CoreError)
		switch {
		case ok && coreErr.Code == core.ErrBadParams:
			return links.NewCreateLinkBadRequest().WithPayload(&models.Error{
				Message: coreErr.Message,
			})
		case ok && coreErr.Code == core.ErrUnauthorized:
			return links.NewCreateLinkUnauthorized().WithPayload(&models.Error{
				Message: "Unauthorized access"})
		case ok && coreErr.Code == core.ErrAccountNotFound:
			return links.NewCreateLinkUnauthorized().WithPayload(&models.Error{
				Message: "Account not found for this username: " + principal.(*core.User).Username})
		case ok && coreErr.Code == core.ErrAccountDisabled:
			return links.NewCreateLinkUnauthorized().WithPayload(&models.Error{
				Message: "Account disabled"})
		default:
			return links.NewCreateLinkInternalServerError().WithPayload(&models.Error{
				Message: "Internal server error",
			})
		}
	}

	return links.NewCreateLinkOK().WithPayload(&models.LinkCreated{
		ShortURL:    link.ShortUrl.String(),
		ShortURLKey: link.ShortUrlKey,
	})
}
