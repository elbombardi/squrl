package routes

import (
	"fmt"

	"github.com/elbombardi/squrl/src/redirection_service/core"
	"github.com/elbombardi/squrl/src/redirection_service/util"
	"github.com/gofiber/fiber/v2"
)

type Routes struct {
	LinksManager core.LinksManager
	*util.Config
}

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
	app.Get("/health", s.Routes.HealthcheckRoute)
	app.Get("/*", s.Routes.RedirectRoute)
	return app.Listen(fmt.Sprintf("%v:%v", s.Host, s.Port))
}
