package dao

import (
	"server-monitor/pkg/server/dal/do"
	"server-monitor/pkg/server/entity/bo"
	"time"
)

type ServerDao struct {
	BaseDao[do.ServerDO, uint]
}

func (dao ServerDao) GetServerInfo(id uint) *bo.ServerInfoBO {
	var Server do.ServerDO
	result := dao.DB().Where("id = ?", id).First(&Server)

	if result.Error != nil {
		return nil
	}

	var serverInfo do.ServerInfoDO
	dao.DB().Where("server_id = ?", id).First(&serverInfo)

	// 获取最后一次状态
	var serverState do.ServerStateDO
	dao.DB().Where("server_id = ?", id).Order("create_time desc").First(&serverState)

	return &bo.ServerInfoBO{
		ServerDO:      &Server,
		ServerInfoDO:  &serverInfo,
		ServerStateDO: &serverState,
	}
}

func (dao ServerDao) GetServerInfoList() []*bo.ServerInfoBO {
	list := dao.GetList()
	if list == nil {
		return nil
	}
	var result []*bo.ServerInfoBO
	for _, v := range list {
		result = append(result, dao.GetServerInfo(v.ID))
	}
	return result
}

func (dao ServerDao) GetServerInfoListByGroupId(groupId uint) []*bo.ServerInfoBO {
	var list []do.ServerDO
	dao.DB().Where("group_id = ?", groupId).Find(&list)

	var result []*bo.ServerInfoBO
	for _, v := range list {
		result = append(result, dao.GetServerInfo(v.ID))
	}
	return result
}

func (dao ServerDao) GetServerGroups() []*bo.ServerGroupBO {
	var groups []do.ServerGroupDO
	dao.DB().Model(&do.ServerGroupDO{}).Order("sort desc").Find(&groups)

	var list []*bo.ServerGroupBO

	for _, v := range groups {
		var result bo.ServerGroupBO

		serverInfos := dao.GetServerInfoListByGroupId(v.ID)
		result.GroupName = v.GroupName
		result.Servers = serverInfos

		list = append(list, &result)
	}

	return list
}

func (dao ServerDao) UpdateLastHeartbeatTime(serverId uint) {
	dao.DB().Model(&do.ServerDO{}).Where("id = ?", serverId).Update("last_online_time", time.Now())
}
