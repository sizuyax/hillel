package http

import (
	"github.com/labstack/echo/v4"
	"hillel/config"
	"hillel/logger"
	"hillel/server/http/routes"
)

type Server struct {
	e    *echo.Echo
	port string
}

func InitWebServer() (*Server, error) {
	e := echo.New()

	cfg, err := config.InitConfig()
	if err != nil {
		logger.Logger.Error(err)
		return nil, err
	}

	logger.NewLogger(*cfg)

	routes.SetupBookRoutes(e)

	return &Server{
		e:    e,
		port: cfg.Port,
	}, nil
}

func (s *Server) StartWebServer() error {
	return s.e.Start(s.port)
}
