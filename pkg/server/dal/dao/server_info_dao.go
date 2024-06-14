package dao

import "server-monitor/pkg/server/dal/do"

type ServerInfoDao struct {
	BaseDao[do.ServerInfoDO, uint]
}

func (dao ServerInfoDao) Save(serverInfoDO *do.ServerInfoDO) {
	serverInfo := dao.GetByServerId(serverInfoDO.ServerId)
	if serverInfo == nil {
		dao.DB().Create(serverInfoDO)
	} else {
		dao.DB().Model(serverInfo).Updates(serverInfoDO)
	}
}

func (dao ServerInfoDao) GetByServerId(serverId uint) *do.ServerInfoDO {
	var serverInfo do.ServerInfoDO
	result := dao.DB().Where("server_id = ?", serverId).First(&serverInfo)
	if result.RowsAffected == 0 {
		return nil
	}
	return &serverInfo
}
