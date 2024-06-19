package bo

import "server-monitor/pkg/server/dal/do"

type ServerFaultBO struct {
	// 服务器名称
	ServerName string `json:"server_name"`
	do.ServerFaultDO
}
