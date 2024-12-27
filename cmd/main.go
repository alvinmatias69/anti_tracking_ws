package main

import (
	"os"

	"github.com/alvinmatias69/anti_tracking_ws/internal/handler"
	"github.com/alvinmatias69/anti_tracking_ws/internal/server"
)

const (
	PORT_ENV_KEY = "PORT"
	DEFAULT_PORT = "8080"
)

func main() {
	var (
		handler = handler.New()
		server  = server.New(handler)
	)

	port, ok := os.LookupEnv(PORT_ENV_KEY)
	if !ok {
		port = DEFAULT_PORT
	}

	server.Start(port)
}
