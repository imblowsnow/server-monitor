package do

import (
	"common/model"
	"time"
)

type ServerState struct {
	ID uint `json:"id" gorm:"primarykey"`
	// 服务器ID
	ServerID uint `json:"server_id" gorm:"index"`
	// 关联服务器信息
	State model.ServerState `json:"state" gorm:"embedded"`
	// 创建时间
	CreatedTime time.Time `json:"created_time" gorm:"autoCreateTime"`
	// 更新时间
	UpdatedTime time.Time `json:"updated_time" gorm:"autoUpdateTime"`
}
