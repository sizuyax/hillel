package main

import (
	"hillel/logger"
	"hillel/server/http"
)

func main() {
	server, err := http.InitWebServer()
	if err != nil {
		logger.Logger.Fatal(err)
	}

	if err := server.StartWebServer(); err != nil {
		logger.Logger.Fatal("failed to start server: ", err)
	}
}
