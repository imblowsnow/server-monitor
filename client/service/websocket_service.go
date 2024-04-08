package service

import (
	"client/utils"
	"common/model"
	"common/websocket"
	websocket2 "github.com/gorilla/websocket"
	"time"
)

type WebsocketService struct {
	*websocket.CommonWebsocketService
}

func NewWebsocketService(conn *websocket2.Conn) *WebsocketService {
	return &WebsocketService{
		CommonWebsocketService: &websocket.CommonWebsocketService{
			Conn: conn,
		},
	}
}

// 开始定时发送服务器状态
func (s *WebsocketService) StartServerStateTicker(seconds time.Duration) {
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
func (s *WebsocketService) SendInit() {
	s.SendMessage(model.MessageTypeInit, model.WebsocketMessageInit{
		Key:        "1234444",
		ServerInfo: utils.GetServerInfo(),
	})
}

// 发送服务器状态
func (s *WebsocketService) sendServerState() {
	s.SendMessage(model.MessageTypeState, utils.GetServerState())
}
