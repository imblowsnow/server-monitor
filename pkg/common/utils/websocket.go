package utils

import (
	"github.com/gorilla/websocket"
	"server-monitor/pkg/common/entity/dto/websocket_message"
)

// 发送消息
func SendWebsocketMessage(conn *websocket.Conn, messageType int, message any) error {
	return conn.WriteJSON(websocket_message.BuildWebsocketMessage(messageType, message))
}

// 回复消息
func ReplyWebsocketMessage(conn *websocket.Conn, messageType int, message any, replyMessageId string) error {
	return conn.WriteJSON(websocket_message.BuildReplyWebsocketMessage(messageType, message, replyMessageId))
}
