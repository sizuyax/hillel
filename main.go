package main

import (
	"flag"
	"hillel/http_server"
	"hillel/tcp_server"
)

func main() {
	var server int

	flag.IntVar(&server, "server", 1, "which one server do you want to start")
	flag.Parse()

	switch server {
	case 1:
		http_server.HttpServer()
	case 2:
		tcp_server.TcpServer()
	}
}
