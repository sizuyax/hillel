package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"project-auction/config"
	"project-auction/database"
	"project-auction/docs"
	"project-auction/lib/logger"
	httpServer "project-auction/server/http"
	"syscall"
	"time"

	"golang.org/x/net/context"
)

//	@title			Project-Auction API
//	@version		1.0
//	@description	Hillel Project
//	@host			http://swagger.io/terms/

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization

//	@BasePath	/
func main() {
	docs.SwaggerInfo.Host = ""

	cfg := config.MustLoad()

	db, err := database.Connect(cfg)
	if err != nil {
		panic(err)
	}

	log := logger.SetupLogger(cfg.LogLevel)
	log = log.With(
		slog.Int("port", cfg.Port),
	)

	router := httpServer.InitWebServer(log, db)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	log.Info("http server started")

	gracefulShutdown(srv, log)
}

func gracefulShutdown(srv *http.Server, log *slog.Logger) {
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
