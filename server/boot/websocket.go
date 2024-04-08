package boot

import (
	websocket2 "common/websocket"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handleUpgradeWebsocket(iw websocket2.IWebSocketEvent, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	websocket2.LoopWebsocketEvent(conn, iw)
}
