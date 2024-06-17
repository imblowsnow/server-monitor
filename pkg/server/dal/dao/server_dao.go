package dao

import (
	"github.com/google/uuid"
	"server-monitor/pkg/common/enum"
	bo2 "server-monitor/pkg/server/common/entity/bo"
	"server-monitor/pkg/server/dal/do"
	"time"
)

type ServerDao struct {
	IBaseDao[do.ServerDO, uint]
}

func NewServerDao() *ServerDao {
	return &ServerDao{
		IBaseDao: &BaseDao[do.ServerDO, uint]{},
	}
}

func (dao ServerDao) GetServerInfo(id uint) *bo2.ServerInfoBO {
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

	return &bo2.ServerInfoBO{
		ServerDO:      &Server,
		ServerInfoDO:  &serverInfo,
		ServerStateDO: &serverState,
	}
}

func (dao ServerDao) GetServerInfoList() []*bo2.ServerInfoBO {
	list := dao.GetList()
	if list == nil {
		return nil
	}
	var result []*bo2.ServerInfoBO
	for _, v := range list {
		result = append(result, dao.GetServerInfo(v.ID))
	}
	return result
}

func (dao ServerDao) GetServerInfoListByGroupId(groupId uint) []*bo2.ServerInfoBO {
	var list []do.ServerDO
	dao.DB().Where("group_id = ?", groupId).Find(&list)

	var result []*bo2.ServerInfoBO
	for _, v := range list {
		result = append(result, dao.GetServerInfo(v.ID))
	}
	return result
}

func (dao ServerDao) GetServerGroups() []*bo2.ServerGroupBO {
	// 获取分组列表
	var serverGroupDao = ServerGroupDao{}
	groups := serverGroupDao.GetList()

	var list []*bo2.ServerGroupBO

	for _, v := range groups {
		var result bo2.ServerGroupBO

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

func (dao ServerDao) Add(server *do.ServerDO) error {
	server.Key = uuid.New().String()
	return dao.IBaseDao.Add(server)
}

func (dao ServerDao) GetStatusNum(status enum.ServerStatus) int64 {
	var count int64
	dao.DB().Model(&do.ServerDO{}).Where("status = ?", status).Count(&count)
	return count
}
