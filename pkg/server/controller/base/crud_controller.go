package base

import (
	"server-monitor/pkg/server/dal/dao"
	"server-monitor/pkg/server/entity"
)

type CrudController[DAO dao.IBaseDao[DO, ID], DO any, ID comparable] struct {
	Dao dao.IBaseDao[DO, ID]
}

func (c CrudController[DAO, DO, ID]) List() []DO {
	return c.Dao.GetList()
}

func (c CrudController[DAO, DO, ID]) Page(page int, size int) (entity.Page[DO], error) {
	var pageE entity.Page[DO]
	pageE.Page = page
	pageE.PageSize = size
	err := c.Dao.Page(&pageE)
	return pageE, err
}

func (c CrudController[DAO, DO, ID]) Get(id ID) *DO {
	return c.Dao.GetById(id)
}

func (c CrudController[DAO, DO, ID]) Create(do *DO) error {
	return c.Dao.Add(do)
}

func (c CrudController[DAO, DO, ID]) Update(id ID, do *DO) error {
	return c.Dao.UpdateById(id, do)
}

func (c CrudController[DAO, DO, ID]) Delete(id ID) error {
	return c.Dao.Delete(id)
}
