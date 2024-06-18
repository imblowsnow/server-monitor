package bo

type MonitorGroupBO struct {
	GroupId   uint               `json:"group_id"`
	GroupName string             `json:"group_name"`
	Servers   []*MonitorServerBO `json:"servers"`
}
