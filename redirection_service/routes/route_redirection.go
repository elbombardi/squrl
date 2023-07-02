package routes

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	db "github.com/elbombardi/squrl/db/sqlc"
	"github.com/elbombardi/squrl/util"
	"github.com/gofiber/fiber/v2"
)

func (r *Routes) RedirectRoute(c *fiber.Ctx) error {
	customerPrefix := c.Params("customer_prefix")
	shortURLKey := c.Params("short_url_key")

	// Retrieve Customer information from the database
	customer, err := r.getCustomerInfo(customerPrefix)
	if err != nil {
		if err == sql.ErrNoRows {
			return page404(c)
		}
		log.Println("Error retrieving Customer information: ", err)
		return page500(c)
	}
	// If the customer is not active, send 404 not found error page
	if customer.Status != "e" {
		log.Printf("Error: Customer %v is disabled \n", customer.Username)
		return page404(c)
	}

	// Retrieve Short URL information from the database
	shortURL, err := r.getShortURLInfo(customer.ID, shortURLKey)
	if err != nil {
		if err == sql.ErrNoRows {
			return page404(c)
		}
		//TODO Send internal error page
		log.Println("Error retrieving short URL information: ", err)
		return page500(c)
	}

	//If the short URL is not active, send 404 not found error page
	if shortURL.Status.String != "e" {
		log.Printf("Error: Short URL /%v is disabled \n", shortURL.ShortUrlKey.String)
		return page404(c)
	}

	// Asynchronously persist click information
	if shortURL.TrackingStatus.String == "e" {
		r.PersistClick(&shortURL, c.IP(), c.Get("User-Agent"))
	}

	// Redirect to the long URL
	return c.Redirect(shortURL.LongUrl, http.StatusFound)
}

func (r *Routes) getCustomerInfo(prefix string) (db.Customer, error) {
	return r.CustomersRepository.GetCustomerByPrefix(context.Background(), prefix)
}

func (r *Routes) getShortURLInfo(customerId int32, key string) (db.ShortUrl, error) {
	return r.ShortURLsRepository.GetShortURLByCustomerIDAndShortURLKey(context.Background(),
		db.GetShortURLByCustomerIDAndShortURLKeyParams{
			CustomerID: customerId,
			ShortUrlKey: sql.NullString{
				String: key,
				Valid:  true,
			},
		})
}

func page404(c *fiber.Ctx) error {
	return c.SendFile(util.ConfigRedirection404Page())
}

func page500(c *fiber.Ctx) error {
	return c.SendFile(util.ConfigRedirection500Page())
}
