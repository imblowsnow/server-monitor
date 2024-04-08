package dao

import (
	"server/entity/do"
)
import "server/config"

type ServerInfoeDao struct {
}

func GetServerInfoDao() ServerInfoeDao {
	return ServerInfoeDao{}
}

// 更新服务器信息
func (s ServerInfoeDao) SaveServerInfo(serverInfo do.ServerInfo) error {
	// 查询是否存在 server_id
	var serverInfoEntity do.ServerInfo
	result := config.Db.Where("server_id = ?", serverInfo.ServerID).Limit(1).Find(&serverInfoEntity)
	if result.Error != nil {
		return result.Error
	}
	if serverInfoEntity.ID != 0 {
		serverInfo.ID = serverInfoEntity.ID
		serverInfo.CreatedTime = serverInfoEntity.CreatedTime
	}
	result = config.Db.Save(&serverInfo)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s ServerInfoeDao) GetServerInfoByServerID(id uint) *do.ServerInfo {
	serverInfo := do.ServerInfo{}
	result := config.Db.Where("server_id = ?", id).Limit(1).Find(&serverInfo)
	if result.Error != nil {
		return nil
	}
	return &serverInfo
}
