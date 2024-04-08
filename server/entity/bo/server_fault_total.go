package bo

import (
	"server/entity/do"
	"time"
)

type ServerFaultTotalBO struct {
	ServerID uint `json:"server_id" gorm:"primarykey"`
	// 名称
	Name string `json:"name" gorm:"type:varchar(100);not null"`
	// 分组
	Group string `json:"group" gorm:"type:varchar(100)"`

	// 统计数据
	Items []*ServerFaultTotalItemBO `json:"items"`
}

type ServerFaultTotalItemBO struct {
	// 时间
	Time string `json:"time"`
	// 服务器总运行时间 单位 秒
	TotalTime time.Duration `json:"total_time"`
	// 服务器故障时间 单位 秒
	FaultTime time.Duration `json:"fault_time"`
	// 故障率
	FaultRate float64 `json:"fault_rate"`

	// 故障列表 前10条
	FaultList []do.ServerFault `json:"fault_list"`
}
