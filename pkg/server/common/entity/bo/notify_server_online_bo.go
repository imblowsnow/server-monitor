package bo

import "server-monitor/pkg/common/entity/dto"

type NotifyServerOnlineBO struct {
	// 服务器ID
	ServerId           uint `json:"server_id"`
	dto.ServerInfoDTO  `json:"server_info"`
	dto.ServerStateDTO `json:"server_state"`
}
