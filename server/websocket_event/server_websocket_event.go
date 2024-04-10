package websocket_event

import (
	"common/model"
	"common/websocket"
	"encoding/json"
	"fmt"
	websocket2 "github.com/gorilla/websocket"
	"server/config"
)

type ServerWebSocketEvent struct {
	websocket.IWebSocketEvent
	conn             *websocket2.Conn
	service          *ServerWebSocketHandle
	FrontNotifyEvent *FrontNotifyEvent
}

func (e *ServerWebSocketEvent) OnConnect(conn *websocket2.Conn) {
	e.conn = conn
	fmt.Println("ServerWebSocketEvent OnConnect", e.conn.RemoteAddr())
}

func (e *ServerWebSocketEvent) OnMessage(message model.WebsocketMessage) {
	fmt.Println("ServerWebSocketEvent OnMessage", e.conn.RemoteAddr(), message)

	if message.MessageType == model.MessageTypeInit {
		initMessage := model.WebsocketMessageInit{}

		err := coverMessage(message.Message, &initMessage)
		if err != nil {
			_ = e.conn.Close()
			fmt.Println("转换json失败", e.conn.RemoteAddr())
			return
		}

		var clientId string
		e.service, clientId, err = NewServerWebSocketHandle(initMessage.Key, e.conn)
		if err != nil {
			_ = e.conn.Close()
			fmt.Println("服务器不存在", initMessage.Key, e.conn.RemoteAddr())
			return
		}

		if config.SocketClients.CheckExistByKey(clientId) {
			clientConn := config.SocketClients.Get(clientId).(*websocket2.Conn)
			_ = clientConn.Close()
			fmt.Println("客户端重复上线，下线上一个客户端", clientConn.RemoteAddr())
		}
		// 添加服务记录
		config.SocketClients.Add(clientId, e.conn)

		e.service.SendMessage(model.MessageTypeInit, nil)
		e.service.SaveServerInfo(initMessage.ServerInfo)
		e.service.CheckFault()
		// TODO 添加一个协程监听器，判断最后一次心跳时间，如果超过一定时间，关闭连接

		// 通知前端web socket
		e.FrontNotifyEvent.OnServerLogin(e.service.serverId, initMessage.ServerInfo)

	} else if !config.SocketClients.CheckExist(e.conn) {
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

		// 通知前端web socket
		e.FrontNotifyEvent.OnServerState(e.service.serverId, serverState)

	}
	// TODO 其他事件的处理

}

func (e *ServerWebSocketEvent) OnClose() {
	fmt.Println("ServerWebSocketEvent OnClose", e.conn.RemoteAddr())
	if e.service != nil {
		e.service.AddFault()
		config.SocketClients.Remove(e.conn)

		// 通知前端web socket
		e.FrontNotifyEvent.OnServerClose(e.service.serverId)
	}
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
