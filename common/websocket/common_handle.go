package websocket

import (
	"common/model"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"time"
)

type CommonWebsocketHandle struct {
	Conn *websocket.Conn
}

func (s *CommonWebsocketHandle) SendMessage(messageType int, data any) error {
	return s.Conn.WriteJSON(model.WebsocketMessage{
		Id:          uuid.New().String(),
		MessageType: messageType,
		Message:     data,
		Time:        time.Now(),
	})
}
