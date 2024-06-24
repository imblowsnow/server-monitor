package main

import (
	"flag"
	"fmt"
	"server-monitor/pkg/client/constants"
	"server-monitor/pkg/client/websocket"
	"time"
)

func main() {
	var host string
	var port int
	var key string
	flag.StringVar(&host, "host", "", "host")
	flag.IntVar(&port, "port", 22251, "port")
	flag.StringVar(&key, "key", "", "key")
	flag.Parse()
	if host == "" {
		fmt.Println("请输入host")
		return
	}
	if key == "" {
		fmt.Println("请输入秘钥")
		return
	}
	// 从参数中获取秘钥
	constants.ServerKey = key
	constants.ServerPort = port
	constants.ServerIP = host

	websocket.LoopWebsocket(host, port)

	ticker := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Println("尝试重新连接：" + host + ":" + fmt.Sprintf("%d", port))
			websocket.LoopWebsocket(host, port)
		}
	}
}
