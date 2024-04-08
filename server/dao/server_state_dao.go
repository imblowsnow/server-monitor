package dao

import (
	"server/entity/do"
)
import "server/config"

type ServerStateDao struct {
}

func GetServerStateDao() ServerStateDao {
	return ServerStateDao{}
}

// 保存服务器状态
func (s ServerStateDao) AddServerState(serverState do.ServerState) error {
	result := config.Db.Create(&serverState)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s ServerStateDao) GetServerStateByServerID(id uint) *do.ServerState {
	serverState := do.ServerState{}
	result := config.Db.Where("server_id = ?", id).Limit(1).Find(&serverState)
	if result.Error != nil {
		return nil
	}
	return &serverState
}
