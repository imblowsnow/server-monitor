package websocket_event

import (
	"client/service"
	"common/model"
	"common/websocket"
	"fmt"
)
import websocket2 "github.com/gorilla/websocket"

type ClientWebSocketEvent struct {
	websocket.IWebSocketEvent
	conn    *websocket2.Conn
	service *service.WebsocketService
}

func (e *ClientWebSocketEvent) OnConnect(conn *websocket2.Conn) {
	e.conn = conn
	fmt.Println("ClientWebSocketEvent OnConnect")
	e.service = service.NewWebsocketService(conn)
	// 发送初始化
	e.service.SendInit()
}

func (e *ClientWebSocketEvent) OnMessage(message model.WebsocketMessage) {
	fmt.Println("ClientWebSocketEvent OnMessage", message)
	// 初始化成功，开始发送心跳
	if message.MessageType == model.MessageTypeInit {
		fmt.Println("开始定时发送服务器状态")
		e.service.StartServerStateTicker(5)
	}
}

func (e *ClientWebSocketEvent) OnClose() {
	fmt.Println("ClientWebSocketEvent OnClose")
}
