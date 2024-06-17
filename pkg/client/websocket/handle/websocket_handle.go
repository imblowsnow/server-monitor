package handle

import (
	"fmt"
	"github.com/gorilla/websocket"
	"server-monitor/pkg/client/utils"
	websocket_message2 "server-monitor/pkg/common/entity/dto/websocket_message"
	"server-monitor/pkg/common/enum"
	commonUtils "server-monitor/pkg/common/utils"
	"time"
)

type WebsocketHandle struct {
}

func (h WebsocketHandle) OnConnected(conn *websocket.Conn) {
	// 发送初始化消息
	message := websocket_message2.MessageServerInitDTO{
		ServerInfoDTO:  utils.GetServerInfo(),
		ServerStateDTO: utils.GetServerState(),
	}
	err := commonUtils.SendWebsocketMessage(conn, enum.MessageServerInit, message)
	if err != nil {
		fmt.Println("发送初始化消息失败:", err)
		conn.Close()
	}
}

func (WebsocketHandle) OnClose(conn *websocket.Conn) {
	fmt.Println("链接关闭")
}

func (h WebsocketHandle) OnServerInitSuccess(conn *websocket.Conn, message websocket_message2.ServerInitSuccessDTO) {
	fmt.Println("初始化成功:", message)
	// 开启心跳 推送服务器状态
	h.startHeartbeat(conn)
}

func (h WebsocketHandle) startHeartbeat(conn *websocket.Conn) {

	go func() {
		// 每分钟推送一次服务器状态
		for {
			select {
			// 30秒推送一次服务器状态
			case <-time.Tick(time.Second * 30):
				fmt.Println("推送服务器状态")
				serverSate := websocket_message2.ServerStateDTO{
					ServerStateDTO: utils.GetServerState(),
				}
				err := commonUtils.SendWebsocketMessage(conn, enum.MessageServerStat, serverSate)
				if err != nil {
					fmt.Println("发送服务器状态失败:", err)
				}
			}
		}
	}()
}
