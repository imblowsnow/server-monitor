package handle

import (
	"fmt"
	"github.com/gorilla/websocket"
	"server-monitor/pkg/common/entity/websocket_message"
	"server-monitor/pkg/common/enum"
	"server-monitor/pkg/common/utils"
)

type ServerWebsocketHandle struct {
}

var serverConnectionManage = NewServerConnectionManage()

func (ServerWebsocketHandle) OnConnected(conn *websocket.Conn) {

}

func (h ServerWebsocketHandle) OnClose(conn *websocket.Conn) {
	serverConnectionManage.RemoveByConn(conn)
}

func (h ServerWebsocketHandle) OnServerInit(conn *websocket.Conn, message websocket_message.ServerInit) {
	// 通过服务器Key获取服务器ID
	serverId := 1

	if serverConnectionManage.Exist(serverId) {
		fmt.Println("服务器已经存在:", serverId)
		// 服务器已经存在
		conn.Close()
		return
	}
	// 保存链接
	serverConnectionManage.Add(serverId, conn)

	serverInitSuccessMessage := websocket_message.ServerInitSuccess{}
	// 发送初始化成功消息
	err := utils.SendWebsocketMessage(conn, enum.ServerMessageInitSuccess, serverInitSuccessMessage)
	if err != nil {
		fmt.Println("发送初始化成功消息失败:", err)
		conn.Close()
	}
}

func (h ServerWebsocketHandle) OnServerStat(conn *websocket.Conn, message websocket_message.ServerState) {
	fmt.Println("收到服务器状态:", message)
}
