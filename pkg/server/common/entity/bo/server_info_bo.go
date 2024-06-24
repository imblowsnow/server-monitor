package bo

import (
	"server-monitor/pkg/common/time"
	"server-monitor/pkg/server/dal/do"
)

type ServerInfoBO struct {
	// 服务器ID
	ID uint `json:"id"`
	// 名称
	Name string `json:"name"`
	// 备注
	Remark string `json:"remark"`
	// 状态 1.正常 0.离线
	Status int    `json:"status"`
	Ip     string `json:"ip"`
	Key    string `json:"key"`
	// 最后在线时间
	LastOnlineTime    time.Time `json:"last_online_time"`
	*do.ServerInfoDO  `json:"server_info"`
	*do.ServerStateDO `json:"server_state"`
}
