package vo

type AddServerVO struct {
	// 组ID
	GroupID uint `json:"group_id"`
	// 服务器名称
	Name string `json:"name"`
	// 服务器备注
	Remark string `json:"remark"`
	// 首页显示 1.显示 0.不显示
	ShowIndex int `json:"show_index" gorm:"type:int"`
}
