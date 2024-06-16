package do

import "time"

type ServerGroupDO struct {
	// ID
	ID uint `json:"id" gorm:"primaryKey;autoIncrement"`
	// 组名
	GroupName string `json:"group_name" gorm:"type:varchar(64)"`
	// 排序
	Sort int `json:"sort" gorm:"type:int;default:1"`
	// 系统 1. 系统值  2. 自定义
	System int `json:"system" gorm:"type:int;default:2"`
	// 创建时间
	CreateTime time.Time `json:"create_time" gorm:"type:datetime;autoCreateTime"`
	// 更新时间
	UpdateTime time.Time `json:"update_time" gorm:"type:datetime;autoUpdateTime"`
}
