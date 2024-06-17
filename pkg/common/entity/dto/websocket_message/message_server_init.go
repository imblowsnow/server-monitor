package websocket_message

import (
	"server-monitor/pkg/common/entity/dto"
)

type MessageServerInitDTO struct {
	// 服务器Key
	Key string `json:"key"`

	dto.ServerInfoDTO `json:"server_info"`

	dto.ServerStateDTO `json:"server_state"`
}
