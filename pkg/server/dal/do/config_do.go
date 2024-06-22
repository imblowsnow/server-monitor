package do

type ConfigDO struct {
	ID int `json:"id" gorm:"primaryKey;autoIncrement"`
	// name
	Name string `json:"name" gorm:"type:varchar(64)"`
	// text
	Value string `json:"value" gorm:"type:text"`
}
