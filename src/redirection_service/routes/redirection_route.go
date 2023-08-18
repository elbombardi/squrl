package routes

import (
	"context"
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/elbombardi/squrl/src/db"
	"github.com/gofiber/fiber/v2"
)

func (r *Routes) RedirectRoute(c *fiber.Ctx) error {
	accountPrefix := c.Params("account_prefix")
	shortURLKey := c.Params("short_url_key")

	// Retrieve Customer information from the database
	account, err := r.getAccountInfo(accountPrefix)
	if err != nil {
		if err == sql.ErrNoRows {
			return page404(c)
		}
		slog.Error("Error retrieving Customer information: ", err)
		return page500(c)
	}
	// If the account is not active, send 404 not found error page
	if !account.Enabled {
		slog.Error("Error: Customer %v is disabled \n", account.Username)
		return page404(c)
	}

	// Retrieve Short URL information from the database
	URL, err := r.getURLInfo(account.ID, shortURLKey)
	if err != nil {
		if err == sql.ErrNoRows {
			return page404(c)
		}
		slog.Error("Error retrieving short URL information: ", err)
		return page500(c)
	}

	//If the short URL is not active, send 404 not found error page
	if !URL.Enabled {
		slog.Error("Error: Short URL /%v is disabled \n", URL.ShortUrlKey.String)
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
			slog.Error("Error inserting new click: ", err)
		}
	}

	// Redirect to the long URL
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
	return c.SendFile("static/404.html")
}

func page500(c *fiber.Ctx) error {
	return c.SendFile("static/500.html")
}
