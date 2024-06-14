package main

import "server-monitor/pkg/client/websocket"

func main() {
	websocket.LoopWebsocket("127.0.0.1", 22251)
}
