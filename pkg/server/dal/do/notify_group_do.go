package do

import "server-monitor/pkg/common/time"

type NotifyGroupDO struct {
	// ID
	ID uint `json:"id" gorm:"primaryKey;autoIncrement"`
	// 通知类型 1.上线 2.下线
	Type int `json:"type" gorm:"type:int"`
	// 名称
	Name string `json:"name" gorm:"type:varchar(64)"`
	// 创建时间
	CreateTime time.Time `json:"create_time" gorm:"type:datetime;autoCreateTime"`
	// 更新时间
	UpdateTime time.Time `json:"update_time" gorm:"type:datetime;autoUpdateTime"`
}
