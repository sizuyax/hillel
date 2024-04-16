package main

import (
	"fmt"
	"golang.org/x/net/context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"project-auction/config"
	"project-auction/docs"
	"project-auction/lib/logger"
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

	log = log.With(
		slog.Int("port", cfg.Port),
	)

	server := httpServer.InitWebServer(log)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: server,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	log.Info("http server started")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	log.Info("received signal to shut down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		panic(err)
	}
}
