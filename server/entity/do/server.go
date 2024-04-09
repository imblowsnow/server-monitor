package do

type Server struct {
	ID uint `json:"id" gorm:"primarykey"`
	// 名称
	Name string `json:"name" gorm:"type:varchar(100);not null"`
	// 备注
	Remark string `json:"remark"`
	// ip地址
	Host string `json:"host" gorm:"type:varchar(128)"`
	// key
	Key string `json:"key" gorm:"type:varchar(100);unique;uniqueIndex"`
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
}
