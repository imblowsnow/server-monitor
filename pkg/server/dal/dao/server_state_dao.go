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
func (dao *ServerStateDao) Add(entity *do.ServerStateDO) error {
	// 获取当前时间 年月日 时分
	addTime := time.Now().Format("2006-01-02 15:04")
	// 判断当前分钟是否已经存在，则不再添加
	var serverState do.ServerStateDO
	dao.DB().Where("server_id = ? and create_time >= ?", entity.ServerId, addTime).First(&serverState)
	if serverState.ID > 0 {
		return nil
	}
	return dao.DB().Create(entity).Error
}
