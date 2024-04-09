package websocket_event

import (
	"client/utils"
	"common/model"
	"common/websocket"
	websocket2 "github.com/gorilla/websocket"
	"time"
)

type ClientWebsocketHandle struct {
	*websocket.CommonWebsocketHandle
}

func NewClientWebsocketHandle(conn *websocket2.Conn) *ClientWebsocketHandle {
	return &ClientWebsocketHandle{
		CommonWebsocketHandle: &websocket.CommonWebsocketHandle{
			Conn: conn,
		},
	}
}

// 开始定时发送服务器状态
func (s *ClientWebsocketHandle) StartServerStateTicker(seconds time.Duration) {
	go func() {
		// 立刻发送一次服务器状态
		s.sendServerState()

		ticker := time.NewTicker(seconds * time.Second)
		for {
			select {
			case <-ticker.C:
				s.sendServerState()
			}
		}
	}()
}

// 发送初始化信息包
func (s *ClientWebsocketHandle) SendInit() {
	s.SendMessage(model.MessageTypeInit, model.WebsocketMessageInit{
		Key:        "1234444",
		ServerInfo: utils.GetServerInfo(),
	})
}

// 发送服务器状态
func (s *ClientWebsocketHandle) sendServerState() {
	s.SendMessage(model.MessageTypeState, utils.GetServerState())
}
