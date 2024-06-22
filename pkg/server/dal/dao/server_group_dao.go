package dao

import (
	"fmt"
	"server-monitor/pkg/server/dal/do"
)

type ServerGroupDao struct {
	IBaseDao[do.ServerGroupDO, uint]
}

func NewServerGroupDao() *ServerGroupDao {
	dao := &ServerGroupDao{
		IBaseDao: &BaseDao[do.ServerGroupDO, uint]{
			Order: "sort desc, id asc",
		},
	}
	dao.Init()
	return dao
}

func (dao *ServerGroupDao) Init() {
	// 初始化系统默认分组
	dao.DB().Where("system = 1").FirstOrCreate(&do.ServerGroupDO{
		GroupName: "默认",
		Sort:      1,
		System:    1,
	})
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
