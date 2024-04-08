package websocket_event

import (
	"common/model"
	websocket "common/websocket"
	"encoding/json"
	"fmt"
	websocket2 "github.com/gorilla/websocket"
	"server/service"
)

type ServerWebSocketEvent struct {
	websocket.IWebSocketEvent
	conn    *websocket2.Conn
	service *service.ServerWebSocketService
}

func (e *ServerWebSocketEvent) OnConnect(conn *websocket2.Conn) {
	e.conn = conn
	fmt.Println("ServerWebSocketEvent OnConnect", e.conn.RemoteAddr())
}

func (e *ServerWebSocketEvent) OnMessage(message model.WebsocketMessage) {
	fmt.Println("ServerWebSocketEvent OnMessage", e.conn.RemoteAddr(), message)

	if message.MessageType == model.MessageTypeInit {
		clientId := "2333"

		initMessage := model.WebsocketMessageInit{}

		err := coverMessage(message.Message, &initMessage)
		if err != nil {
			_ = e.conn.Close()
			fmt.Println("转换json失败", e.conn.RemoteAddr())
			return
		}
		e.service = service.NewServerWebSocketService(clientId, e.conn)

		if e.service.CheckExistById(clientId) {
			clientConn := e.service.GetClient(clientId)
			_ = clientConn.Close()
			fmt.Println("客户端重复上线，下线上一个客户端", clientConn.RemoteAddr())
		}
		e.service.AddClient(clientId, e.conn)
		e.service.SendMessage(model.MessageTypeInit, nil)
		e.service.SaveServerInfo(initMessage.ServerInfo)
	} else if !e.service.CheckExist(e.conn) {
		_ = e.conn.Close()
		fmt.Println("客户端未初始化成功，关闭连接", e.conn.RemoteAddr())
	} else if message.MessageType == model.MessageTypeState {
		serverState := model.ServerState{}
		err := coverMessage(message.Message, &serverState)
		if err != nil {
			fmt.Println("转换json失败", e.conn.RemoteAddr())
			return
		}
		e.service.SaveServerState(serverState)
	}
	// TODO 其他事件的处理

}

func (e *ServerWebSocketEvent) OnClose() {
	fmt.Println("ServerWebSocketEvent OnClose", e.conn.RemoteAddr())
	e.service.RemoveClient(e.conn)
}

func coverMessage(message any, data any) error {
	jsonStr, err := json.Marshal(message)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonStr, data)
	if err != nil {
		return err
	}

	return nil
}
