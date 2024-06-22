package dao

import (
	"server-monitor/pkg/server/dal/do"
	"time"
)

type ServerStateDao struct {
	IBaseDao[do.ServerStateDO, uint]
}

func NewServerStateDao() *ServerStateDao {
	return &ServerStateDao{
		IBaseDao: &BaseDao[do.ServerStateDO, uint]{},
	}
}

func (dao *ServerStateDao) Record(entity *do.ServerStateDO) error {
	if dao.CheckNowMinuteExist(entity.ServerId) {
		return nil
	}
	return dao.DB().Create(entity).Error
}

func (dao *ServerStateDao) CheckNowMinuteExist(serverId uint) bool {
	// 获取当前时间 年月日 时分
	addTime := time.Now().Format("2006-01-02 15:04") + ":00"
	// 判断当前分钟是否已经存在，则不再添加
	var count int64
	dao.DB().Model(&do.ServerStateDO{}).Where("server_id = ? and create_time >= ?", serverId, addTime).Count(&count)
	return count > 0
}
