package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"server-monitor/pkg/client/websocket/event"
	"server-monitor/pkg/common/inner_websocket"
)

func LoopWebsocket(host string, port int) error {
	// 链接服务端websocket
	serverUrl := fmt.Sprintf("ws://%s:%d/ws/client", host, port)
	conn, _, err := websocket.DefaultDialer.Dial(serverUrl, nil)
	if err != nil {
		fmt.Println("连接服务端失败:", err)
		return err
	}

	// 循环监听消息
	inner_websocket.LoopWebsocket(conn, &event.WebsocketEvent{})

	return nil
}
