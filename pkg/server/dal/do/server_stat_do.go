package do

import (
	"server-monitor/pkg/common/entity/bo"
	"time"
)

type ServerStatDO struct {
	// ID
	ID uint `json:"id" gorm:"primaryKey;autoIncrement"`
	// 服务器ID
	ServerId uint `json:"server_id"`

	// 服务器当前信息
	bo.ServerStateBO `gorm:"embedded"`

	// 创建时间
	CreateTime time.Time `json:"create_time" gorm:"type:datetime"`
	// 更新时间
	UpdateTime time.Time `json:"update_time" gorm:"type:datetime"`
}
