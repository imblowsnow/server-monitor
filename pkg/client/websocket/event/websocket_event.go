package event

import (
	"fmt"
	"github.com/gorilla/websocket"
	"server-monitor/pkg/client/websocket/handle"
	"server-monitor/pkg/common/entity/websocket_message"
	"server-monitor/pkg/common/enum"
	"server-monitor/pkg/common/inner_websocket"
)

type WebsocketEvent struct {
	inner_websocket.WebsocketEvent
}

var websocketHandle = handle.WebsocketHandle{}

func (WebsocketEvent) OnConnected(conn *websocket.Conn) {
	websocketHandle.OnConnected(conn)
}

// 链接关闭
func (WebsocketEvent) OnClose(conn *websocket.Conn) {
	websocketHandle.OnClose(conn)
}

// 收到消息
func (WebsocketEvent) OnMessage(conn *websocket.Conn, websocketMessage websocket_message.WebsocketMessage) {
	fmt.Println("收到服务端消息:", websocketMessage)
	switch websocketMessage.MessageType {
	case enum.ServerMessageInitSuccess:
		message, err := websocket_message.ToMessageType(websocketMessage, websocket_message.ServerInitSuccess{})
		if err != nil {
			fmt.Println("解析消息失败:", err)
			return
		}
		websocketHandle.OnServerInitSuccess(conn, message)
		break
	}
}
