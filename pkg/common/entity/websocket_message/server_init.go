package websocket_message

import (
	"server-monitor/pkg/common/entity/bo"
)

type ServerInit struct {
	// 服务器Key
	Key string `json:"key"`

	bo.ServerInfoBO

	bo.ServerStateBO
}
