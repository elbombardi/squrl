package routes

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

/*
Route that handles health check requests (GET /health)
*/
func (r *Routes) HealthcheckRoute(c *fiber.Ctx) error {
	c.Status(http.StatusOK)
	return c.SendString("OK")
}
