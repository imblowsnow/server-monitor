package websocket_message

import "server-monitor/pkg/common/model"

type ServerInit struct {
	// 服务器Key
	Key string `json:"key"`

	model.ServerInfo

	model.ServerState
}
