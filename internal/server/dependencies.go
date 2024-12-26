package server

import "github.com/lxzan/gws"

type handler interface {
	OnOpen(socket *gws.Conn)                          // connection is established
	OnClose(socket *gws.Conn, err error)              // received a close frame or input/output error occurs
	OnPing(socket *gws.Conn, payload []byte)          // received a ping frame
	OnPong(socket *gws.Conn, payload []byte)          // received a pong frame
	OnMessage(socket *gws.Conn, message *gws.Message) // received a text/binary frame
}
