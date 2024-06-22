package event

import (
	"fmt"
	"github.com/gorilla/websocket"
	websocket_message2 "server-monitor/pkg/common/entity/dto/websocket_message"
	"server-monitor/pkg/common/inner_websocket"
)

type WebClientWebsocketEvent struct {
	inner_websocket.WebsocketEvent
}

func (WebClientWebsocketEvent) OnConnected(conn *websocket.Conn) {
	fmt.Println("web client websocket connected")
	webClientWebsocketHandle.OnClientConnected(conn)
}

func (WebClientWebsocketEvent) OnClose(conn *websocket.Conn) {
	fmt.Println("web client websocket close")
	webClientWebsocketHandle.OnClientClose(conn)
}

// 收到消息
func (WebClientWebsocketEvent) OnMessage(conn *websocket.Conn, websocketMessage websocket_message2.WebsocketMessageDTO) {

}
