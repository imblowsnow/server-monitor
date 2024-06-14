package bo

import "server-monitor/pkg/server/dal/do"

type ServerInfoBO struct {
	do.ServerDO      `json:"server"`
	*do.ServerInfoDO `json:"server_info"`
	*do.ServerStatDO `json:"server_state"`
}
