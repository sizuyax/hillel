package http

import (
	"github.com/labstack/echo/v4"
	"hillel/config"
	"hillel/handlers"
	"hillel/logger"
)

type Server struct {
	e    *echo.Echo
	Port string
}

func InitWebServer() (*Server, error) {
	e := echo.New()

	cfg, err := config.InitConfig()
	if err != nil {
		logger.Logger.Error(err)
		return nil, err
	}

	logger.NewLogger(*cfg)

	handlers.AllHandlers(e)

	return &Server{
		e:    e,
		Port: cfg.Port,
	}, nil
}

func (s *Server) StartWebServer() error {
	return s.e.Start(s.Port)
}
