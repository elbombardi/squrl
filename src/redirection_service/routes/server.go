package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	Port   int
	Host   string
	Routes *Routes
}

func NewServer(port int, host string, routes *Routes) *Server {
	return &Server{
		Port:   port,
		Host:   host,
		Routes: routes,
	}
}

func (s *Server) Serve() error {
	app := fiber.New()
	app.Get("/:account_prefix/:short_url_key", s.Routes.RedirectRoute)
	return app.Listen(fmt.Sprintf("%v:%v", s.Host, s.Port))
}
