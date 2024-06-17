package utils

import (
	"github.com/gorilla/websocket"
	"server-monitor/pkg/common/entity/dto/websocket_message"
)

func SendWebsocketMessage(conn *websocket.Conn, success int, message any) error {
	return conn.WriteJSON(websocket_message.BuildWebsocketMessage(success, message))
}
