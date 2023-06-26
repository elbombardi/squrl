package routes

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	db "github.com/elbombardi/squrl/db/sqlc"
	"github.com/gofiber/fiber/v2"
)

func (r *Routes) RedirectRoute(c *fiber.Ctx) error {
	customerPrefix := c.Params("customer_prefix")
	shortURLKey := c.Params("short_url_key")

	// Retrieve Customer information from the database
	customer, err := r.getCustomerInfo(customerPrefix)
	if err != nil {
		if err == sql.ErrNoRows {
			//TODO Send 404 not found error page
			return c.SendStatus(http.StatusNotFound)
		}
		//TODO Send internal error page
		log.Println("Error retrieving Customer information: ", err)
		return c.SendStatus(http.StatusInternalServerError)
	}
	// If the customer is not active, send 404 not found error page
	if customer.Status != "e" {
		//TODO Send 404 not found error page
		return c.SendStatus(http.StatusNotFound)
	}

	// Retrieve Short URL information from the database
	shortURL, err := r.getShortURLInfo(customer.ID, shortURLKey)
	if err != nil {
		if err == sql.ErrNoRows {
			//TODO Send 404 not found error page
			return c.SendStatus(http.StatusNotFound)
		}
		//TODO Send internal error page
		log.Println("Error retrieving short URL information: ", err)
		return c.SendStatus(http.StatusInternalServerError)
	}

	//If the short URL is not active, send 404 not found error page
	if shortURL.Status.String != "e" {
		//TODO Send 404 not found error page
		return c.SendStatus(http.StatusNotFound)
	}

	// Asynchronously persist click information
	if shortURL.TrackingStatus.String == "e" {
		r.PersistClick(&shortURL)
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
