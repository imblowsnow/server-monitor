package websocket_event

import (
	"common/model"
	"common/websocket"
	"fmt"
	websocket2 "github.com/gorilla/websocket"
	"server/dao"
	"server/entity/do"
	"server/service"
	"time"
)

var cacheServerWebSocketHandle = &CacheWebsocketHandle{}
var serverService = service.ServerService{}
var serverStateService = service.ServerStateService{}
var serverInfoService = service.ServerInfoService{}
var serverFaultDao = dao.ServerFaultDao{}

type ServerWebSocketHandle struct {
	*CacheWebsocketHandle
	*websocket.CommonWebsocketHandle
	serverId uint
}

func NewServerWebSocketHandle(key string, conn *websocket2.Conn) (*ServerWebSocketHandle, string, error) {
	server, err := serverService.GetServerByKey(key)
	if err != nil {
		return nil, "", err
	}
	if server.ID == 0 {
		return nil, "", err
	}
	serverId := fmt.Sprintf("%d", server.ID)
	return &ServerWebSocketHandle{
		serverId: server.ID,
		CommonWebsocketHandle: &websocket.CommonWebsocketHandle{
			Conn: conn,
		},
		CacheWebsocketHandle: cacheServerWebSocketHandle,
	}, serverId, nil
}

// 保存服务器状态
func (s *ServerWebSocketHandle) SaveServerState(state model.ServerState) {
	fmt.Println("SaveServerState", state)
	serverStateService.AddServerState(do.ServerState{
		ServerID: s.serverId,
		State:    state,
	})
}

// 保存服务器信息
func (s *ServerWebSocketHandle) SaveServerInfo(info model.ServerInfo) {
	fmt.Println("SaveServerInfo", info)
	serverInfoService.SaveServerInfo(do.ServerInfo{
		ServerID: s.serverId,
		Info:     info,
	})
}

// 新增故障 离线后调用
func (s *ServerWebSocketHandle) AddFault() {
	serverFaultDao.AddServerFault(&do.ServerFault{
		ServerID:  s.serverId,
		StartTime: time.Now(),
	})
}

// 检查故障，更新故障时间，如果故障小于5分钟，不处理
func (s *ServerWebSocketHandle) CheckFault() {
	serverFault := serverFaultDao.GetUndoServerFaultByServerID(s.serverId)
	if serverFault.ServerID == 0 {
		return
	}
	if time.Now().Sub(serverFault.StartTime) < 5*time.Minute {
		serverFaultDao.ClearServerFault(s.serverId)
		return
	}
	serverFault.EndTime = time.Now()
	serverFault.Duration = int(serverFault.EndTime.Sub(serverFault.StartTime).Seconds())
	serverFaultDao.UpdateServerFault(&serverFault)
}
