package bo

type ServerGroupBO struct {
	GroupName string          `json:"group_name"`
	Servers   []*ServerInfoBO `json:"servers"`
}
