package websocket_event

import (
	"common/model"
	"common/websocket"
	websocket2 "github.com/gorilla/websocket"
)

type WebWebSocketEvent struct {
	websocket.IWebSocketEvent
	conn *websocket2.Conn
}

func (e *WebWebSocketEvent) OnConnect(conn *websocket2.Conn) {
	e.conn = conn
}

func (e *WebWebSocketEvent) OnMessage(message model.WebsocketMessage) {

}

func (e *WebWebSocketEvent) OnClose() {

}
