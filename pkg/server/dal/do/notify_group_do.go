package do

type NotifyGroupDO struct {
	// ID
	ID uint `json:"id" gorm:"primaryKey;autoIncrement"`
	// 通知类型 1.上线 2.下线
	Type int `json:"type" gorm:"type:int"`
	// 名称
	Name string `json:"name" gorm:"type:varchar(64)"`
	// 创建时间
	CreateTime string `json:"create_time" gorm:"type:datetime"`
	// 更新时间
	UpdateTime string `json:"update_time" gorm:"type:datetime"`
}
