package inner_websocket

import (
	"github.com/gorilla/websocket"
	"server-monitor/pkg/common/entity/dto/websocket_message"
)

type WebsocketEvent interface {
	// 链接成功
	OnConnected(conn *websocket.Conn)

	// 链接关闭
	OnClose(conn *websocket.Conn)

	// 收到消息
	OnMessage(conn *websocket.Conn, websocketMessage websocket_message.WebsocketMessageDTO)
}
