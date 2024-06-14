package websocket_message

import "server-monitor/pkg/common/entity"

type ServerInit struct {
	// 服务器Key
	Key string `json:"key"`

	entity.ServerInfo

	entity.ServerState
}
