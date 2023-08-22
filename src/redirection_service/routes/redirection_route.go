package routes

import (
	"net/http"

	"github.com/elbombardi/squrl/src/redirection_service/core"
	"github.com/gofiber/fiber/v2"
)

// Route that handles redirections requests (GET "/:account_prefix/:short_url_key")
func (r *Routes) RedirectRoute(c *fiber.Ctx) error {
	accountPrefix := c.Params("account_prefix")
	shortURLKey := c.Params("short_url_key")

	link, err := r.LinksManager.Resolve(&core.ResolveLinkParams{
		ShortUrl:      c.Request().URI().String(),
		AccountPrefix: accountPrefix,
		ShortUrlKey:   shortURLKey,
		UserAgent:     c.Get("User-Agent"),
		IpAddress:     c.IP(),
	})
	if err != nil {
		coreError, ok := err.(core.CoreError)
		switch {
		case ok && coreError.Code == core.ErrAccountNotFound:
			fallthrough
		case ok && coreError.Code == core.ErrAccountDisabled:
			fallthrough
		case ok && coreError.Code == core.ErrLinkNotFound:
			fallthrough
		case ok && coreError.Code == core.ErrLinkDisabled:
			return page404(c)
		case ok && coreError.Code == core.ErrBadParams:
			fallthrough
		default:
			return page500(c)
		}
	}

	// Redirect to the long URL
	return c.Redirect(link.LongUrl.String(), http.StatusFound)
}

func page404(c *fiber.Ctx) error {
	c.Response().Header.SetContentType(fiber.MIMETextHTML)
	c.Response().SetStatusCode(http.StatusNotFound)
	return c.SendString(RESPONSE_404)
}

func page500(c *fiber.Ctx) error {
	c.Response().Header.SetContentType(fiber.MIMETextHTML)
	c.Response().SetStatusCode(http.StatusInternalServerError)
	return c.SendString(RESPONSE_500)
}
