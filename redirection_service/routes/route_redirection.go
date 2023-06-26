package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Customer struct {
	ID     int
	Prefix string
	Status string
}

type ShortURL struct {
	ID     int
	Key    string
	Status string
	URL    string
}

func (r *Routes) RedirectRoute(c *fiber.Ctx) error {
	log.Println("Redirecting..")
	customerPrefix := c.Params("customer_prefix")
	shortURLKey := c.Params("short_url_key")

	// Retrieve Customer information from the database
	customer, err := getCustomerInfo(customerPrefix)
	if err != nil || customer == nil || customer.Status != "enabled" {
		return c.SendStatus(http.StatusNotFound)
	}

	// Retrieve Short URL information from the database
	shortURL, err := getShortURLInfo(shortURLKey)
	if err != nil || shortURL == nil || shortURL.Status != "enabled" {
		return c.SendStatus(http.StatusNotFound)
	}

	// Asynchronously persist click information
	go persistClickInfo(customer.ID, shortURL.ID)

	// Redirect to the long URL
	return c.Redirect(shortURL.URL, http.StatusFound)
}

func getCustomerInfo(prefix string) (*Customer, error) {
	// Simulate retrieving Customer information from the database
	// Replace this with your actual database query
	// Return nil if Customer not found
	// Replace this with your error handling logic
	return &Customer{
		ID:     1,
		Prefix: prefix,
		Status: "enabled",
	}, nil
}

func getShortURLInfo(key string) (*ShortURL, error) {
	// Simulate retrieving Short URL information from the database
	// Replace this with your actual database query
	// Return nil if Short URL not found
	// Replace this with your error handling logic
	return &ShortURL{
		ID:     1,
		Key:    key,
		Status: "enabled",
		URL:    "https://example.com",
	}, nil
}

func persistClickInfo(customerID, shortURLID int) {
	// Simulate persisting click information in the database
	// Replace this with your actual database logic
	fmt.Printf("Persisting click information for Customer ID %d and Short URL ID %d\n", customerID, shortURLID)
}
