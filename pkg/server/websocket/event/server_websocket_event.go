package event

import (
	"fmt"
	"github.com/gorilla/websocket"
	websocket_message2 "server-monitor/pkg/common/entity/dto/websocket_message"
	"server-monitor/pkg/common/enum"
	"server-monitor/pkg/common/inner_websocket"
	"server-monitor/pkg/server/websocket/handle"
)

type ServerWebsocketEvent struct {
	inner_websocket.WebsocketEvent
}

var serverWebsocketHandle = handle.ServerWebsocketHandle{}

func (ServerWebsocketEvent) OnConnected(conn *websocket.Conn) {
	serverWebsocketHandle.OnConnected(conn)
}

func (ServerWebsocketEvent) OnClose(conn *websocket.Conn) {
	serverWebsocketHandle.OnClose(conn)
}

// 收到消息
func (ServerWebsocketEvent) OnMessage(conn *websocket.Conn, websocketMessage websocket_message2.WebsocketMessageDTO) {
	fmt.Println("收到客户端消息:", websocketMessage)
	switch websocketMessage.MessageType {
	case enum.MessageServerInit:
		serverInitMessage, err := websocket_message2.ToMessageType(websocketMessage, websocket_message2.MessageServerInitDTO{})
		if err != nil {
			fmt.Println("解析消息失败:", err)
			return
		}
		serverWebsocketHandle.OnServerInit(conn, serverInitMessage)
		break
	case enum.MessageServerStat:
		serverStatMessage, err := websocket_message2.ToMessageType(websocketMessage, websocket_message2.ServerStateDTO{})
		if err != nil {
			fmt.Println("解析消息失败:", err)
			return
		}
		serverWebsocketHandle.OnServerState(conn, serverStatMessage)
	}
}
