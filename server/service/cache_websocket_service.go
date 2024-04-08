package service

import "github.com/gorilla/websocket"

type CacheWebsocketService struct {
	clientMap map[string]*websocket.Conn
}

func (s *CacheWebsocketService) GetClientId(conn *websocket.Conn) string {
	if s.clientMap == nil {
		return ""
	}
	for key, tmpConn := range s.clientMap {
		if tmpConn == conn {
			return key
		}
	}
	return ""
}
func (s *CacheWebsocketService) GetClient(clientId string) *websocket.Conn {
	if s.clientMap == nil {
		return nil
	}
	return s.clientMap[clientId]
}
func (s *CacheWebsocketService) AddClient(clientId string, conn *websocket.Conn) {
	if s.clientMap == nil {
		s.clientMap = make(map[string]*websocket.Conn)
	}
	s.clientMap[clientId] = conn
}

func (s *CacheWebsocketService) RemoveClient(conn *websocket.Conn) {
	if s.clientMap == nil {
		s.clientMap = make(map[string]*websocket.Conn)
	}
	clientId := s.GetClientId(conn)

	delete(s.clientMap, clientId)
}

// 检查是否授权了
func (s *CacheWebsocketService) CheckExist(conn *websocket.Conn) bool {
	if s.clientMap == nil {
		return false
	}
	clientId := s.GetClientId(conn)
	return clientId != ""
}

func (s *CacheWebsocketService) CheckExistById(clientId string) bool {
	if s.clientMap == nil {
		return false
	}
	return s.clientMap[clientId] != nil
}
