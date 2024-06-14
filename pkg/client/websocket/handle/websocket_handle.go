package handle

import (
	"fmt"
	"github.com/gorilla/websocket"
	"server-monitor/pkg/client/utils"
	"server-monitor/pkg/common/entity/websocket_message"
	"server-monitor/pkg/common/enum"
	"time"
)

type WebsocketHandle struct {
}

func (h WebsocketHandle) OnConnected(conn *websocket.Conn) {
	// 发送初始化消息
	message := websocket_message.ServerInit{
		ServerInfo:  utils.GetServerInfo(),
		ServerState: utils.GetServerState(),
	}
	err := h.sendMessage(conn, enum.MessageServerInit, message)
	if err != nil {
		fmt.Println("发送初始化消息失败:", err)
		conn.Close()
	}
}

func (WebsocketHandle) OnClose(conn *websocket.Conn) {
	fmt.Println("链接关闭")
}

func (h WebsocketHandle) OnServerInitSuccess(conn *websocket.Conn, message websocket_message.ServerInitSuccess) {
	fmt.Println("初始化成功:", message)
	// 开启心跳 推送服务器状态
	h.startHeartbeat(conn)
}

func (h WebsocketHandle) startHeartbeat(conn *websocket.Conn) {
	go func() {
		// 每分钟推送一次服务器状态
		for {
			select {
			case <-time.Tick(time.Minute):
				serverSate := websocket_message.ServerState{
					ServerState: utils.GetServerState(),
				}
				err := h.sendMessage(conn, enum.MessageServerStat, serverSate)
				if err != nil {
					fmt.Println("发送服务器状态失败:", err)
				}
			}
		}
	}()
}

func (WebsocketHandle) sendMessage(conn *websocket.Conn, messageType int, message interface{}) error {
	return conn.WriteJSON(websocket_message.BuildWebsocketMessage(messageType, message))
}
