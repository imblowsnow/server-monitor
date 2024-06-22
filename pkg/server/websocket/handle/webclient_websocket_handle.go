package handle

import (
	"github.com/gorilla/websocket"
	websocket_message2 "server-monitor/pkg/common/entity/dto/websocket_message"
	"server-monitor/pkg/common/enum"
	time2 "server-monitor/pkg/common/time"
	"server-monitor/pkg/common/utils"
	"server-monitor/pkg/server/common/entity/bo"
	"server-monitor/pkg/server/dal/do"
	"time"
)

type WebClientWebsocketHandle struct {
	clients map[*websocket.Conn]bool
}

func NewWebClientWebsocketHandle() WebClientWebsocketHandle {
	return WebClientWebsocketHandle{
		clients: make(map[*websocket.Conn]bool),
	}
}

func (h *WebClientWebsocketHandle) OnClientConnected(conn *websocket.Conn) {
	h.clients[conn] = true
}

func (h *WebClientWebsocketHandle) OnClientClose(conn *websocket.Conn) {
	delete(h.clients, conn)
}

func (w WebClientWebsocketHandle) NotifyServerState(serverState do.ServerStateDO) {
	if time.Time(serverState.UpdateTime).IsZero() {
		serverState.UpdateTime = time2.Time(time.Now())
	}
	w.notifyAll(enum.WebMessageServerState, serverState)
}

func (w WebClientWebsocketHandle) NotifyServerOnline(serverId uint, message websocket_message2.MessageServerInitDTO) {
	notifyServerOnlineBO := bo.NotifyServerOnlineBO{
		ServerId:       serverId,
		ServerInfoDTO:  message.ServerInfoDTO,
		ServerStateDTO: message.ServerStateDTO,
	}
	w.notifyAll(enum.WebMessageServerOnline, notifyServerOnlineBO)
}

func (w WebClientWebsocketHandle) NotifyServerOffline(serverId uint) {
	w.notifyAll(enum.WebMessageServerOffline, serverId)
}

func (w *WebClientWebsocketHandle) notifyAll(eventType int, message interface{}) {
	for conn, _ := range w.clients {
		_ = utils.SendWebsocketMessage(conn, eventType, message)
	}
}
