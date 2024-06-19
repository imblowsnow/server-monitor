package main

import (
	"fmt"
	"server-monitor/pkg/client/websocket"
	"time"
)

func main() {
	websocket.LoopWebsocket("127.0.0.1", 22251)

	ticker := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Println("尝试重新连接")
			websocket.LoopWebsocket("127.0.0.1", 22251)
		}
	}
}
