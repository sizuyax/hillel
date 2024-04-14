package http

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"project-auction/server/http/routes"
)

type Server struct {
	e    *echo.Echo
	log  *logrus.Logger
	port int
}

func InitWebServer(port int, log *logrus.Logger) *Server {
	e := echo.New()

	routes.SetupRoutes(e)

	return &Server{
		e:    e,
		log:  log,
		port: port,
	}
}

func (s Server) StartWebServer() error {
	return s.e.Start(fmt.Sprintf(":%v", s.port))
}

func (s Server) StopWebServer(ctx context.Context) error {
	return s.e.Shutdown(ctx)
}
