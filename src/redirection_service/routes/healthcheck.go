package routes

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (r *Routes) HealthcheckRoute(c *fiber.Ctx) error {
	c.Status(http.StatusOK)
	return c.SendString("OK")
}
