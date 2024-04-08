package service

import (
	"common/model"
	websocket2 "common/websocket"
	"fmt"
	"github.com/gorilla/websocket"
	"server/entity/do"
	"time"
)

var cacheServerWebsocketService = &CacheWebsocketService{}

type ServerWebSocketService struct {
	*CacheWebsocketService
	*websocket2.CommonWebsocketService
	serverId uint
}

func NewServerWebSocketService(key string, conn *websocket.Conn) (*ServerWebSocketService, string, error) {
	server, err := serverDao.GetServerByKey(key)
	if err != nil {
		return nil, "", err
	}
	if server.ID == 0 {
		return nil, "", err
	}
	serverId := fmt.Sprintf("%d", server.ID)
	return &ServerWebSocketService{
		serverId: server.ID,
		CommonWebsocketService: &websocket2.CommonWebsocketService{
			Conn: conn,
		},
		CacheWebsocketService: cacheServerWebsocketService,
	}, serverId, nil
}

// 保存服务器状态
func (s *ServerWebSocketService) SaveServerState(state model.ServerState) {
	fmt.Println("SaveServerState", state)
	serverStateDao.AddServerState(do.ServerState{
		ServerID: s.serverId,
		State:    state,
	})
}

// 保存服务器信息
func (s *ServerWebSocketService) SaveServerInfo(info model.ServerInfo) {
	fmt.Println("SaveServerInfo", info)
	serverInfoDao.SaveServerInfo(do.ServerInfo{
		ServerID: s.serverId,
		Info:     info,
	})
}

// 新增故障 离线后调用
func (s *ServerWebSocketService) AddFault() {
	serverFaultDao.AddServerFault(&do.ServerFault{
		ServerID:  s.serverId,
		StartTime: time.Now(),
	})
}

// 检查故障，更新故障时间，如果故障小于5分钟，不处理
func (s *ServerWebSocketService) CheckFault() {
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
