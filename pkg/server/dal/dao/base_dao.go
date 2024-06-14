package dao

import (
	"fmt"
	"server-monitor/pkg/server/constants"
	"server-monitor/pkg/server/entity"
)

type BaseDao[T any] struct {
}

func (BaseDao[T]) GetList() []T {
	list := []T{}
	result := constants.DB.Find(&list)
	if result.Error != nil {
		fmt.Println("查询列表失败", result.Error)
		return nil
	}
	return list
}

func (BaseDao[T]) Page(page *entity.Page[T]) {
	result := constants.DB.Model(page.List).Count(&page.TotalCount)
	if result.Error != nil {
		fmt.Println("查询总数失败", result.Error)
		return
	}
	result = constants.DB.Offset((page.Page - 1) * page.PageSize).Limit(page.PageSize).Find(&page.List)
	if result.Error != nil {
		fmt.Println("查询列表失败", result.Error)
		return
	}
}

func (BaseDao[T]) Delete(id int) bool {
	var entity T
	result := constants.DB.Delete(entity, id)
	if result.Error != nil {
		fmt.Println("删除失败", id, result.Error)
		return false
	}
	return true
}

func (BaseDao[T]) Add(entity T) bool {
	result := constants.DB.Create(entity)
	if result.Error != nil {
		fmt.Println("新增失败", entity, result.Error)
		return false
	}
	return true
}

func (BaseDao[T]) Update(entity T) bool {
	result := constants.DB.Save(entity)
	if result.Error != nil {
		fmt.Println("更新失败", entity, result.Error)
		return false
	}
	return true
}

func (BaseDao[T]) GetById(id int) T {
	var entity T
	result := constants.DB.First(&entity, id)
	if result.Error != nil {
		fmt.Println("查询失败", id, result.Error)
		return entity
	}
	return entity
}
