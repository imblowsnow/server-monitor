package dao

import (
	"fmt"
	"gorm.io/gorm"
	"server-monitor/pkg/server/config"
	"server-monitor/pkg/server/entity"
)

// DO 数据库实体
// ID 主键类型
type BaseDao[DO any, ID comparable] struct {
	_     IBaseDao[DO, ID]
	Order string
}

func (dao *BaseDao[DO, ID]) GetList() []DO {
	list := []DO{}

	query := dao.DB()
	if dao.Order != "" {
		query = query.Order(dao.Order)
	}
	result := query.Find(&list)
	if result.Error != nil {
		fmt.Println("查询列表失败", result.Error)
		return nil
	}
	return list
}

func (dao *BaseDao[DO, ID]) Page(page *entity.Page[DO]) error {
	// 默认页数值
	if page.Page == 0 {
		page.Page = 1
	}
	if page.PageSize == 0 {
		page.PageSize = 10
	}
	result := dao.DB().Model(page.List).Count(&page.TotalCount)
	if result.Error != nil {
		fmt.Println("查询总数失败", result.Error)
		return result.Error
	}
	page.TotalPage = int((page.TotalCount + int64(page.PageSize) - 1) / int64(page.PageSize))

	query := dao.DB()
	if page.Order != "" {
		query = query.Order(page.Order)
	} else if dao.Order != "" {
		query = query.Order(dao.Order)
	}
	result = query.Offset((page.Page - 1) * page.PageSize).Limit(page.PageSize).Find(&page.List)
	if result.Error != nil {
		fmt.Println("查询列表失败", result.Error)
		return result.Error
	}
	return nil
}

func (dao *BaseDao[DO, ID]) Delete(id ID) error {
	var entity DO
	result := dao.DB().Delete(entity, id)
	if result.Error != nil {
		fmt.Println("删除失败", id, result.Error)
		return result.Error
	}
	return nil
}

func (dao *BaseDao[DO, ID]) GetById(id ID) *DO {
	var entity DO
	result := dao.DB().First(&entity, id)
	if result.Error != nil {
		fmt.Println("查询失败", id, result.Error)
		return nil
	}
	if result.RowsAffected == 0 {
		return nil
	}
	return &entity
}

func (dao *BaseDao[DO, ID]) Add(entity *DO) error {
	result := dao.DB().Create(&entity)
	if result.Error != nil {
		fmt.Println("新增失败", entity, result.Error)
		return result.Error
	}
	return nil
}

func (dao *BaseDao[DO, ID]) UpdateById(id ID, entity *DO) error {
	result := dao.DB().Model(entity).Where("id = ?", id).Updates(entity)
	if result.Error != nil {
		fmt.Println("更新失败", entity, result.Error)
		return result.Error
	}
	return nil
}

func (dao *BaseDao[DO, ID]) Update(entity *DO) error {
	result := dao.DB().Save(&entity)
	if result.Error != nil {
		fmt.Println("更新失败", entity, result.Error)
		return result.Error
	}
	return nil
}

func (*BaseDao[DO, ID]) DB() *gorm.DB {
	return config.DB()
}
