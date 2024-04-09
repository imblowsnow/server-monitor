package boot

import (
	"client/websocket_event"
	"common/websocket"
	"fmt"
	websocket2 "github.com/gorilla/websocket"
)

func Run(host string, port int) error {
	// 链接 websocket 服务
	// format ws://%s:%d/client
	url := fmt.Sprintf("ws://%s:%d/client", host, port)

	conn, _, err := websocket2.DefaultDialer.Dial(url, nil)
	if err != nil {
		return err
	}

	// 开始 websocket 事件循环
	websocket.LoopWebsocketEvent(conn, websocket_event.NewClientWebSocketEvent("1234444"))

	return nil
}
