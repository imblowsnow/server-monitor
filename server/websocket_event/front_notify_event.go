package websocket_event

import (
	"common/model"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"server/utils"
	"time"
)

var cacheFrontWebSockets = &utils.Cache{}

// 用于处理 服务器相关上线下线事件
type FrontNotifyEvent struct {
}

func (e *FrontNotifyEvent) OnServerLogin(id uint, info model.ServerInfo) {
	serverInfo := serverService.GetServerInfoById(id)
	if serverInfo == nil {
		fmt.Println("服务器不存在", id)
		return
	}
	e.SendAll(model.FrontMessageServerLogin, serverInfo)
}

func (e *FrontNotifyEvent) OnServerState(id uint, serverState model.ServerState) {
	e.SendAll(model.FrontMessageServerState, map[string]any{
		"id":    id,
		"state": serverState,
	})
}

func (e *FrontNotifyEvent) OnServerClose(id uint) {
	e.SendAll(model.FrontMessageServerLogout, id)
}

func (e *FrontNotifyEvent) SendAll(messageType int, data any) {
	if cacheFrontWebSockets.GetAll() == nil {
		return
	}
	for _, conn := range cacheFrontWebSockets.GetAll() {
		_ = conn.(*websocket.Conn).WriteJSON(model.WebsocketMessage{
			Id:          uuid.New().String(),
			MessageType: messageType,
			Message:     data,
			Time:        time.Now(),
		})
	}
}
