package bo

import (
	"common/model"
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
	CreatedTime int64 `json:"created_time" gorm:"autoCreateTime"`
	// 更新时间
	UpdatedTime int64 `json:"updated_time" gorm:"autoUpdateTime"`
	// 服务器详细信息
	Info *model.ServerInfo `json:"info"`
	// 服务器最后状态信息
	State *model.ServerState `json:"state"`

	// 是否存活
	Live bool `json:"live"`
}
