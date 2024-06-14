package do

import "time"

type ServerDO struct {
	// ID
	ID uint `json:"id" gorm:"primaryKey;autoIncrement"`
	// 组ID
	GroupID uint `json:"group_id" gorm:"type:int;index"`
	// 密钥
	Key string `json:"key" gorm:"type:varchar(64);index"`
	// 名称
	Name string `json:"name" gorm:"type:varchar(64)"`
	// IP
	IP string `json:"ip" gorm:"type:varchar(64)"`
	// 备注
	Remark string `json:"remark" gorm:"type:varchar(64)"`
	// 首页显示 1.显示 0.不显示
	ShowIndex int `json:"show_index" gorm:"type:int"`
	// 状态 1.正常 0.离线
	Status int `json:"status" gorm:"type:int"`
	// 最后在线时间
	LastOnlineTime time.Time `json:"last_online_time" gorm:"type:datetime"`
	// 创建时间
	CreateTime time.Time `json:"create_time" gorm:"type:datetime;autoCreateTime"`
	// 更新时间
	UpdateTime time.Time `json:"update_time" gorm:"type:datetime;autoUpdateTime"`
}
