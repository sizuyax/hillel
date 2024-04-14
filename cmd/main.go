package main

import (
	"golang.org/x/net/context"
	"os"
	"os/signal"
	"project-auction/config"
	"project-auction/docs"
	"project-auction/logger"
	"project-auction/server/http"
	"syscall"
	"time"
)

// @title  			Project-Auction API
// @version			1.0
// @description 	Hillel Project
// @host 			http://swagger.io/terms/

// @BasePath 		/
func main() {
	docs.SwaggerInfo.Host = ""

	cfg := config.MustLoad()

	log := logger.SetupLogger(cfg.LogLevel)

	server := http.InitWebServer(cfg.Port, log)

	go func() {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

		<-stop

		log.Info("received signal to shut down the server")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.StopWebServer(ctx); err != nil {
			log.Errorf("error shutting down server: %v", err)
		} else {
			log.Info("server gracefully stopped")
		}
	}()

	if err := server.StartWebServer(); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
