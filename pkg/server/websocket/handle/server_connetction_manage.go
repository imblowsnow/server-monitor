package handle

import (
	"github.com/gorilla/websocket"
	"sync"
)

type ServerConnectionManage struct {
	mutex sync.Mutex
	// serverId to conn
	serverConns map[int]*websocket.Conn
	// conn to serverId
	connServers map[*websocket.Conn]int
}

func NewServerConnectionManage() *ServerConnectionManage {
	return &ServerConnectionManage{
		serverConns: make(map[int]*websocket.Conn),
		connServers: make(map[*websocket.Conn]int),
	}
}

func (m *ServerConnectionManage) Add(serverId int, conn *websocket.Conn) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.serverConns[serverId] = conn
	m.connServers[conn] = serverId
}

func (m *ServerConnectionManage) Remove(serverId int) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	conn := m.serverConns[serverId]
	delete(m.serverConns, serverId)
	delete(m.connServers, conn)
}

func (m *ServerConnectionManage) RemoveByConn(conn *websocket.Conn) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	serverId := m.connServers[conn]
	delete(m.serverConns, serverId)
	delete(m.connServers, conn)
}

func (m *ServerConnectionManage) GetConn(serverId int) *websocket.Conn {
	return m.serverConns[serverId]
}

func (m *ServerConnectionManage) GetServerId(conn *websocket.Conn) int {
	return m.connServers[conn]
}

func (m *ServerConnectionManage) Exist(serverId int) bool {
	_, ok := m.serverConns[serverId]
	return ok
}
