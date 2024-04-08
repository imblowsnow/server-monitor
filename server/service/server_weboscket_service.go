package service

import (
	"common/model"
	websocket2 "common/websocket"
	"fmt"
	"github.com/gorilla/websocket"
)

var cacheServerWebsocketService = &CacheWebsocketService{}

type ServerWebSocketService struct {
	*CacheWebsocketService
	*websocket2.CommonWebsocketService
	clientId string
}

func NewServerWebSocketService(clientId string, conn *websocket.Conn) *ServerWebSocketService {
	return &ServerWebSocketService{
		clientId: clientId,
		CommonWebsocketService: &websocket2.CommonWebsocketService{
			Conn: conn,
		},
		CacheWebsocketService: cacheServerWebsocketService,
	}
}

// 保存服务器状态
func (s ServerWebSocketService) SaveServerState(state model.ServerState) {
	fmt.Println("SaveServerState", state)
}

// 保存服务器信息
func (s ServerWebSocketService) SaveServerInfo(info model.ServerInfo) {
	fmt.Println("SaveServerInfo", info)
}
