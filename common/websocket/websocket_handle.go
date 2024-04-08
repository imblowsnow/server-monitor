package websocket

import (
	"common/model"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/gorilla/websocket"
	"strings"
)

func LoopWebsocketEvent(conn *websocket.Conn, iw IWebSocketEvent) {
	defer conn.Close()

	iw.OnConnect(conn)

	conn.SetCloseHandler(func(code int, text string) error {
		iw.OnClose()
		return nil
	})

	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			// 判断conn是否关闭
			if strings.Contains(err.Error(), "closed") || websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				_ = conn.CloseHandler()(websocket.CloseGoingAway, err.Error())
			}
			fmt.Println(err)
			return
		}

		websocketMessage := model.WebsocketMessage{}
		err = json.Unmarshal(data, &websocketMessage)
		if err != nil {
			fmt.Println("websocket收到消息错误：", err, string(data))
			return
		}
		iw.OnMessage(websocketMessage)
	}
}
