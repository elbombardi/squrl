package routes

import (
	"context"
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/elbombardi/squrl/src/db"
	"github.com/gofiber/fiber/v2"
)

/*
Route that handles redirections requests (GET "/:account_prefix/:short_url_key")
*/
func (r *Routes) RedirectRoute(c *fiber.Ctx) error {
	accountPrefix := c.Params("account_prefix")
	shortURLKey := c.Params("short_url_key")

	// Retrieve Customer information from the database
	account, err := r.getAccountInfo(accountPrefix)
	if err != nil {
		if err == sql.ErrNoRows {
			slog.Error("Account not found", "Account", accountPrefix)
			return page404(c)
		}
		slog.Error("Unexpected error while retrieving account information", "Details", err)
		return page500(c)
	}
	// If the account is not active, send 404 not found error page
	if !account.Enabled {
		slog.Info("Click on a URL belonging to a disabled account", "URL", c.Request().URI())
		return page404(c)
	}

	// Retrieve Short URL information from the database
	URL, err := r.getURLInfo(account.ID, shortURLKey)
	if err != nil {
		if err == sql.ErrNoRows {
			return page404(c)
		}
		slog.Error("Unexpected error retrieving short URL information", "Details", err)
		return page500(c)
	}

	//If the short URL is not active, send 404 not found error page
	if !URL.Enabled {
		slog.Info("Click on a disabled URL", "URL", c.Request().URI())
		return page404(c)
	}

	// Persist click information
	if URL.TrackingEnabled {
		err = r.ClickRepository.InsertNewClick(context.Background(), db.InsertNewClickParams{
			UrlID:     URL.ID,
			UserAgent: sql.NullString{String: c.Get("User-Agent"), Valid: true},
			IpAddress: sql.NullString{String: c.IP(), Valid: true},
		})
		if err != nil {
			slog.Error("Unexptected error while inserting a new click", "Details", err)
		}
	}

	// Redirect to the long URL
	slog.Info("URL Clicked", "URL", c.Request().URI())
	return c.Redirect(URL.LongUrl, http.StatusFound)
}

func (r *Routes) getAccountInfo(prefix string) (db.Account, error) {
	return r.AccountRepository.GetAccountByPrefix(context.Background(), prefix)
}

func (r *Routes) getURLInfo(accountId int32, key string) (db.Url, error) {
	return r.URLRepository.GetURLByAccountIDAndShortURLKey(context.Background(),
		db.GetURLByAccountIDAndShortURLKeyParams{
			AccountID: accountId,
			ShortUrlKey: sql.NullString{
				String: key,
				Valid:  true,
			},
		},
	)
}

func page404(c *fiber.Ctx) error {
	c.Response().Header.SetContentType(fiber.MIMETextHTML)
	return c.SendString(RESPONSE_404)
}

func page500(c *fiber.Ctx) error {
	c.Response().Header.SetContentType(fiber.MIMETextHTML)
	return c.SendString(RESPONSE_500)
}
