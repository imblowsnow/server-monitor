package bo

import (
	"common/model"
	"time"
)

type ServerStatisticsBO struct {
	ServerID uint `json:"id" gorm:"primarykey"`
	// 名称
	Name string `json:"name" gorm:"type:varchar(100);not null"`
	// 分组
	Group string `json:"group" gorm:"type:varchar(100)"`
	// 服务器详细信息
	Info *model.ServerInfo `json:"info"`
	// 统计数据
	Items []*ServerStatisticsItemBO `json:"items"`
}

type ServerStatisticsItemBO struct {
	// 时间
	Time time.Time `json:"time"`
	// 服务器状态
	State *model.ServerState `json:"state"`
}
