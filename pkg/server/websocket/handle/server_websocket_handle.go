package handle

import (
	"fmt"
	"github.com/gorilla/websocket"
	websocket_message2 "server-monitor/pkg/common/entity/dto/websocket_message"
	"server-monitor/pkg/common/enum"
	time2 "server-monitor/pkg/common/time"
	"server-monitor/pkg/common/utils"
	"server-monitor/pkg/server/dal/dao"
	"server-monitor/pkg/server/dal/do"
	"strings"
	"time"
)

type ServerWebsocketHandle struct {
	webClientWebsocketHandle *WebClientWebsocketHandle
}

var serverConnectionManage = NewServerConnectionManage()
var serverDao = dao.NewServerDao()
var serverStateDao = dao.NewServerStateDao()
var serverInfoDao = dao.NewServerInfoDao()
var serverFaultDao = dao.NewServerFaultDao()

func NewServerWebsocketHandle(webClientWebsocketHandle *WebClientWebsocketHandle) ServerWebsocketHandle {
	return ServerWebsocketHandle{
		webClientWebsocketHandle: webClientWebsocketHandle,
	}
}

func (ServerWebsocketHandle) OnConnected(conn *websocket.Conn) {

}

// 客户端关闭事件
func (h ServerWebsocketHandle) OnClose(conn *websocket.Conn) {
	serverId := serverConnectionManage.GetServerId(conn)
	if serverId == 0 {
		return
	}
	fmt.Println("服务器关闭:", serverId)
	serverConnectionManage.RemoveByConn(conn)

	// 通知服务器状态离线
	h.notifyServerOffline(conn, uint(serverId))
}

// 客户端初始化事件
func (h ServerWebsocketHandle) OnServerInit(conn *websocket.Conn, message websocket_message2.MessageServerInitDTO) {
	// 通过服务器Key获取服务器ID
	serverId := int(serverDao.GetServerIdByKey(message.Key))

	if serverId == 0 {
		fmt.Println("秘钥错误:", message.Key)
		conn.Close()
		return
	}

	if serverConnectionManage.Exist(serverId) {
		fmt.Println("服务器已经存在:", serverId)
		// 服务器已经存在
		conn.Close()
		return
	}
	// 保存链接
	serverConnectionManage.Add(serverId, conn)

	serverInitSuccessMessage := websocket_message2.ServerInitSuccessDTO{}
	// 发送初始化成功消息
	err := utils.SendWebsocketMessage(conn, enum.ServerMessageInitSuccess, serverInitSuccessMessage)
	if err != nil {
		fmt.Println("发送初始化成功消息失败:", err)
		conn.Close()
		return
	}

	h.notifyServerOnline(conn, uint(serverId), message)
}

// 客户端状态事件
func (h ServerWebsocketHandle) OnServerState(conn *websocket.Conn, message websocket_message2.ServerStateDTO) {
	fmt.Println("收到服务器状态:", message)
	serverId := serverConnectionManage.GetServerId(conn)

	// 通知服务器状态
	h.notifyServerState(conn, do.ServerStateDO{
		ServerId:       uint(serverId),
		ServerStateDTO: message.ServerStateDTO,
	})
}

func (h ServerWebsocketHandle) notifyServerState(conn *websocket.Conn, serverState do.ServerStateDO) {
	serverId := serverState.ServerId
	// 记录服务器状态
	serverStateDao.Record(&serverState)
	// 更新服务器最后心跳时间
	serverDao.UpdateLastHeartbeatTime(serverId)

	// 推送web 客户端
	h.webClientWebsocketHandle.NotifyServerState(serverState)
}

func (h ServerWebsocketHandle) notifyServerOnline(conn *websocket.Conn, serverId uint, message websocket_message2.MessageServerInitDTO) {

	// 恢复服务器故障
	serverFaultDao.RecoverFault(serverId)

	ip := conn.RemoteAddr().String()
	// 只保留IP ，去掉端口
	ip = ip[:strings.LastIndex(ip, ":")]
	serverDao.UpdateById(serverId, &do.ServerDO{
		Status:         int(enum.ServerStatusOnline),
		LastOnlineTime: time2.Time(time.Now()),
		IP:             ip,
	})

	// 更新服务器信息
	serverInfoDao.Save(&do.ServerInfoDO{
		ServerId:      serverId,
		ServerInfoDTO: message.ServerInfoDTO,
	})

	// 通知服务器状态
	h.notifyServerState(conn, do.ServerStateDO{
		ServerId:       serverId,
		ServerStateDTO: message.ServerStateDTO,
	})

	// 推送web 客户端
	h.webClientWebsocketHandle.NotifyServerOnline(serverId, message)
}

func (h ServerWebsocketHandle) notifyServerOffline(conn *websocket.Conn, serverId uint) {
	// 更新服务器状态为离线
	serverDao.UpdateStatus(serverId, enum.ServerStatusOffline)

	// 记录服务器掉线
	serverFaultDao.RecordFault(serverId, "服务器掉线")

	// 推送到web客户端
	h.webClientWebsocketHandle.NotifyServerOffline(serverId)
}
