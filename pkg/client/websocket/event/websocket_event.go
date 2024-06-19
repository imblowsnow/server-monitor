package event

import (
	"fmt"
	"github.com/gorilla/websocket"
	"server-monitor/pkg/client/websocket/handle"
	websocket_message2 "server-monitor/pkg/common/entity/dto/websocket_message"
	"server-monitor/pkg/common/enum"
	"server-monitor/pkg/common/inner_websocket"
)

type WebsocketEvent struct {
	inner_websocket.WebsocketEvent
}

var websocketHandle = handle.NewWebsocketHandle()

func (WebsocketEvent) OnConnected(conn *websocket.Conn) {
	websocketHandle.OnConnected(conn)
}

// 链接关闭
func (WebsocketEvent) OnClose(conn *websocket.Conn) {
	websocketHandle.OnClose(conn)
}

// 收到消息
func (WebsocketEvent) OnMessage(conn *websocket.Conn, websocketMessage websocket_message2.WebsocketMessageDTO) {
	fmt.Println("收到服务端消息:", websocketMessage)
	switch websocketMessage.MessageType {
	case enum.ServerMessageInitSuccess:
		message, err := websocket_message2.ToMessageType(websocketMessage, websocket_message2.ServerInitSuccessDTO{})
		if err != nil {
			fmt.Println("解析消息失败:", err)
			return
		}
		websocketHandle.OnServerInitSuccess(conn, message)
		break
	}
}
