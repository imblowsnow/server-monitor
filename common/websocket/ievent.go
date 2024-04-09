package websocket

import (
	"common/model"
	websocket2 "github.com/gorilla/websocket"
)

type IWebSocketEvent interface {
	OnConnect(conn *websocket2.Conn)
	OnMessage(model.WebsocketMessage)
	OnClose()
}
