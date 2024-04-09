package dao

import (
	"server/config"
	"server/entity/do"
)

type ServerFaultDao struct {
}

// 获取单个故障信息
func (s *ServerFaultDao) GetServerFaultByServerID(serverId uint) *do.ServerFault {

	serverFault := do.ServerFault{}

	config.Db.Where("server_id = ?", serverId).First(&serverFault)

	return &serverFault
}

func (s *ServerFaultDao) GetUndoServerFaultByServerID(serverId uint) do.ServerFault {

	serverFault := do.ServerFault{}

	config.Db.Where("server_id = ?", serverId).Where("end_time is null").Order("start_time desc").Limit(1).Find(&serverFault)

	return serverFault
}

// 清除未处理的故障
func (s *ServerFaultDao) ClearServerFault(serverId uint) {
	config.Db.Where("server_id = ?", serverId).Where("end_time is null").Delete(&do.ServerFault{})
}

// 新增故障
func (s *ServerFaultDao) AddServerFault(serverFault *do.ServerFault) {
	config.Db.Create(serverFault)
}

// 更新故障
func (s *ServerFaultDao) UpdateServerFault(serverFault *do.ServerFault) {
	config.Db.Save(serverFault)
}

// 时间段查询故障列表
func (s *ServerFaultDao) GetServerFaultListByTime(serverId uint, startTime, endTime int64) []do.ServerFault {

	serverFaultList := make([]do.ServerFault, 0)

	config.Db.Raw("select * from server_fault where server_id = ? and start_time >= ? and (end_time < ? or end_time is null)", serverId, startTime, endTime).Scan(&serverFaultList)

	return serverFaultList
}
