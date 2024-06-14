package dao

import (
	"fmt"
	"server-monitor/pkg/server/dal/do"
	"server-monitor/pkg/server/entity"
	"testing"
)

func TestServerDao(t *testing.T) {
	serverDao := ServerDao{}

	page := entity.Page[do.ServerDO]{
		Page:     1,
		PageSize: 10,
	}
	serverDao.Page(&page)

	data := serverDao.GetById(1)

	fmt.Println("查询列表", page, data)
}