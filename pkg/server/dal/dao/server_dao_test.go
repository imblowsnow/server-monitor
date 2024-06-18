package dao

import (
	"fmt"
	"server-monitor/pkg/server/common/entity"
	"server-monitor/pkg/server/dal/do"
	"testing"
	"time"
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
func TestGetMonitorServerStatisticsList(t *testing.T) {
	serverDao := NewServerDao()
	list := serverDao.GetMonitorServerStatisticsList(1, time.Hour, 24)
	fmt.Println("查询列表", list)
}
