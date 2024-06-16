package dao

import (
	"fmt"
	"server-monitor/pkg/server/dal/do"
)

type ServerGroupDao struct {
	BaseDao[do.ServerGroupDO, uint]
}

func NewServerGroupDao() *ServerGroupDao {
	return &ServerGroupDao{
		BaseDao: BaseDao[do.ServerGroupDO, uint]{
			Order: "sort desc, id asc",
		},
	}
}

func (dao *ServerGroupDao) Delete(id uint) error {
	entity := dao.GetById(id)
	if entity == nil {
		return fmt.Errorf("id不存在")
	}
	if entity.System == 1 {
		return fmt.Errorf("系统默认分组不能删除")
	}
	result := dao.DB().Delete(entity, id)
	if result.Error != nil {
		fmt.Println("删除失败", id, result.Error)
		return result.Error
	}
	return nil
}
