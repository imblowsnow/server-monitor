package websocket_event

import (
	"common/model"
	"common/websocket"
	"github.com/google/uuid"
	websocket2 "github.com/gorilla/websocket"
)

type FrontWebSocketEvent struct {
	websocket.IWebSocketEvent
	conn *websocket2.Conn

	service *FrontWebSocketHandle
}

func (e *FrontWebSocketEvent) OnConnect(conn *websocket2.Conn) {
	e.conn = conn
	e.service = NewFrontWebSocketHandle()
	cacheFrontWebSockets.Add(uuid.New().String(), e.conn)
}

func (e *FrontWebSocketEvent) OnMessage(message model.WebsocketMessage) {

}

func (e *FrontWebSocketEvent) OnClose() {
	cacheFrontWebSockets.Remove(e.conn)
}
