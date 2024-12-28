package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lxzan/gws"
)

type Server struct {
	handler gws.Event
}

func New(handler handler) *Server {
	return &Server{
		handler: handler,
	}
}

func (s *Server) Start(port string) {
	upgrader := gws.NewUpgrader(s.handler, &gws.ServerOption{
		ParallelEnabled:   true,
		Recovery:          gws.Recovery,
		PermessageDeflate: gws.PermessageDeflate{Enabled: true},
	})

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	http.HandleFunc("/connect", func(writer http.ResponseWriter, req *http.Request) {
		socket, err := upgrader.Upgrade(writer, req)
		if err != nil {
			return
		}
		go func() {
			socket.ReadLoop()
		}()
	})

	log.Printf("starting websocket server on port: %s", port)

	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
