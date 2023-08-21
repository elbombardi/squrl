package handlers

import (
	"github.com/elbombardi/squrl/src/api_service/api/models"
	"github.com/elbombardi/squrl/src/api_service/api/operations/links"
	"github.com/elbombardi/squrl/src/api_service/core"
	"github.com/go-openapi/runtime/middleware"
)

// Hanlder for the PUT /urls endpoint
func (h *Handlers) HandleUpdateLink(params links.UpdateLinkParams, principal any) middleware.Responder {

	if params.Body == nil {
		return links.NewUpdateLinkBadRequest().WithPayload(&models.Error{
			Message: "Request body is required",
		})
	}

	if params.Body.Status != "" &&
		params.Body.Status != "active" &&
		params.Body.Status != "inactive" {
		return links.NewUpdateLinkBadRequest().WithPayload(&models.Error{
			Message: "Invalid status, should be one of the two values: 'active', 'inactive'",
		})
	}

	if params.Body.TrackingStatus != "" &&
		params.Body.TrackingStatus != "active" &&
		params.Body.TrackingStatus != "inactive" {
		return links.NewUpdateLinkBadRequest().WithPayload(&models.Error{
			Message: "Invalid tracking status, should be one of the two values: 'active', 'inactive'",
		})
	}

	encodeStatus(params.Body.Status)

	link, err := h.LinksManager.Update(&core.LinkUpdateParams{
		ShortUrlKey: params.Body.ShortURLKey,
		NewLongURL: core.Optional[string]{
			Value: params.Body.NewLongURL,
			IsSet: params.Body.NewLongURL != "",
		},
		Enabled:         encodeStatus(params.Body.Status),
		TrackingEnabled: encodeStatus(params.Body.TrackingStatus),
	}, principal.(*core.User))

	if err != nil {
		coreError, ok := err.(core.CoreError)
		switch {
		case ok && coreError.Code == core.ErrBadParams:
			return links.NewUpdateLinkBadRequest().WithPayload(&models.Error{
				Message: coreError.Message,
			})
		case ok && coreError.Code == core.ErrUnauthorized:
			return links.NewUpdateLinkUnauthorized().WithPayload(&models.Error{
				Message: "Unauthorized access"})
		case ok && coreError.Code == core.ErrAccountDisabled:
			return links.NewUpdateLinkUnauthorized().WithPayload(&models.Error{
				Message: "Account disabled"})
		case ok && coreError.Code == core.ErrAccountNotFound:
			return links.NewUpdateLinkUnauthorized().WithPayload(&models.Error{
				Message: "Account not found. Username: " + principal.(*core.User).Username})
		case ok && coreError.Code == core.ErrLinkNotFound:
			return links.NewUpdateLinkNotFound().WithPayload(&models.Error{
				Message: "Link not found. Username: " + principal.(*core.User).Username + ", ShortUrlKey: " + params.Body.ShortURLKey})
		default:
			return links.NewUpdateLinkInternalServerError().WithPayload(&models.Error{
				Message: "Internal server error",
			})
		}
	}

	return links.NewUpdateLinkOK().WithPayload(&models.LinkUpdated{
		LongURL:        link.LongUrl.String(),
		Status:         decodeStatus(link.Enabled),
		TrackingStatus: decodeStatus(link.TrackingEnabled),
	})
}
