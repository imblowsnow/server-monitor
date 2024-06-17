package inner_websocket

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"server-monitor/pkg/common/entity/dto/websocket_message"
)

func LoopWebsocket(conn *websocket.Conn, event WebsocketEvent) {
	defer conn.Close()

	event.OnConnected(conn)

	conn.SetCloseHandler(func(code int, text string) error {
		event.OnClose(conn)
		return nil
	})

	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			event.OnClose(conn)
			return
		}

		websocketMessage := websocket_message.WebsocketMessageDTO{}
		err = json.Unmarshal(data, &websocketMessage)
		if err != nil {
			fmt.Println("解析消息失败:", err)
			continue
		}

		event.OnMessage(conn, websocketMessage)
	}
}
