package handler

import (
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/lxzan/gws"
)

const (
	userIDKey = "user_id"
)

type Handler struct {
	connectedClients map[string]*gws.Conn
}

func New() *Handler {
	return &Handler{
		connectedClients: map[string]*gws.Conn{},
	}
}

func (h *Handler) OnOpen(socket *gws.Conn) {
	userID := uuid.New().String()
	socket.Session().Store(userIDKey, userID)

	err := socket.WriteMessage(gws.OpcodeText, []byte(fmt.Sprintf("success connecting with id: %v", userID)))
	if err != nil {
		log.Printf("error while sending message: %v", err)
		return
	}

	h.connectedClients[userID] = socket
	log.Printf("client connected with id: %s", userID)
}

func (h *Handler) OnClose(socket *gws.Conn, err error) {
	userID, err := getUserId(socket)
	if err != nil {
		log.Printf("error while closing connection: %v", err)
		return
	}

	delete(h.connectedClients, userID)
	log.Printf("client disconnected with id: %s", userID)
}

func (h *Handler) OnPing(socket *gws.Conn, payload []byte) {
	_ = socket.WritePong(nil)
}

func (h *Handler) OnPong(socket *gws.Conn, payload []byte) {}

func (h *Handler) OnMessage(socket *gws.Conn, message *gws.Message) {
	defer message.Close()

	userID, err := getUserId(socket)
	if err != nil {
		log.Printf("error getting user id: %v", err)
		return
	}

	broadcaster := gws.NewBroadcaster(gws.OpcodeText, message.Bytes())
	defer broadcaster.Close()

	for receiverID, conn := range h.connectedClients {
		if receiverID == userID {
			continue
		}

		err = broadcaster.Broadcast(conn)
		if err != nil {
			log.Printf("error while sending message to %s with error: %v", receiverID, err)
		}
	}
}

func getUserId(socket *gws.Conn) (string, error) {
	userID, ok := socket.Session().Load(userIDKey)
	if !ok {
		return "", errors.New("user id not found in session")
	}

	userIDStr, ok := userID.(string)
	if !ok {
		return "", errors.New("malformed user id (not string)")
	}

	return userIDStr, nil
}
