package do

import (
	"common/model"
)

type ServerInfo struct {
	ID       uint `json:"id" gorm:"primarykey"`
	ServerID uint `json:"server_id" gorm:"uniqueIndex"`
	// 关联服务器信息
	Info model.ServerInfo `json:"info" gorm:"embedded"`
	// 创建时间
	CreatedTime int64 `json:"created_time" gorm:"autoCreateTime"`
	// 更新时间
	UpdatedTime int64 `json:"updated_time" gorm:"autoUpdateTime"`
}
