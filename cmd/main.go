package main

import (
	"github.com/alvinmatias69/anti_tracking_ws/internal/handler"
	"github.com/alvinmatias69/anti_tracking_ws/internal/server"
)

func main() {
	var (
		handler = handler.New()
		server  = server.New(handler)
	)

	server.Start(":8080")
}
