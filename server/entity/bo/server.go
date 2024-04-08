package bo

import (
	"server/entity/do"
	"time"
)

type ServerBO struct {
	ID uint `json:"id" gorm:"primarykey"`
	// 名称
	Name string `json:"name" gorm:"type:varchar(100);not null"`
	// 分组
	Group string `json:"group" gorm:"type:varchar(100)"`
	// 版本号
	Version string `json:"version" gorm:"type:varchar(100)"`
	// 游客隐藏
	Hide int `json:"hide" gorm:"type:int(1);default:0"`
	// 排序
	Index int `json:"index" gorm:"type:int(10);default:0"`
	// 创建时间
	CreatedTime time.Time `json:"created_time" gorm:"autoCreateTime"`
	// 更新时间
	UpdatedTime time.Time `json:"updated_time" gorm:"autoUpdateTime"`
	// 服务器详细信息
	Info *do.ServerInfo `json:"info"`
	// 服务器最后状态信息
	State *do.ServerState `json:"state"`
}
