package base

import (
	"server-monitor/pkg/server/dal/dao"
	"server-monitor/pkg/server/entity"
)

type ICrudController[DAO dao.IBaseDao[DO, ID], DO any, ID comparable] interface {
	List() []DO

	Page(page int, size int) (entity.Page[DO], error)

	Get(id ID) *DO

	Create(do *DO) error

	Update(id ID, do *DO) error

	Delete(id ID) error
}
