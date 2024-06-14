package handle

import (
	"fmt"
	"github.com/gorilla/websocket"
	"server-monitor/pkg/common/entity/websocket_message"
	"server-monitor/pkg/common/enum"
	"server-monitor/pkg/common/utils"
	"server-monitor/pkg/server/dal/dao"
	"server-monitor/pkg/server/dal/do"
)

type ServerWebsocketHandle struct {
}

var serverConnectionManage = NewServerConnectionManage()
var serverDao = dao.ServerDao{}
var serverStateDao = dao.ServerStateDao{}
var serverInfoDao = dao.ServerInfoDao{}

func (ServerWebsocketHandle) OnConnected(conn *websocket.Conn) {

}

// 客户端关闭事件
func (h ServerWebsocketHandle) OnClose(conn *websocket.Conn) {
	serverId := serverConnectionManage.GetServerId(conn)
	fmt.Println("服务器关闭:", serverId)
	serverConnectionManage.RemoveByConn(conn)
}

// 客户端初始化事件
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

	// 更新服务器最后心跳时间

	serverInfoDao.Save(&do.ServerInfoDO{
		ServerId:   uint(serverId),
		ServerInfo: message.ServerInfo,
	})
	serverStateDao.Add(&do.ServerStateDO{
		ServerId:    uint(serverId),
		ServerState: message.ServerState,
	})

	serverDao.UpdateLastHeartbeatTime(uint(serverId))
}

// 客户端状态事件
func (h ServerWebsocketHandle) OnServerState(conn *websocket.Conn, message websocket_message.ServerState) {
	fmt.Println("收到服务器状态:", message)
	serverId := serverConnectionManage.GetServerId(conn)
	serverStateDao.Add(&do.ServerStateDO{
		ServerId:    uint(serverId),
		ServerState: message.ServerState,
	})

	serverDao.UpdateLastHeartbeatTime(uint(serverId))
}
