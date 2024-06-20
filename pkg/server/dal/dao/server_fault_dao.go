package dao

import (
	"fmt"
	time2 "server-monitor/pkg/common/time"
	"server-monitor/pkg/server/dal/do"
	"time"
)

type ServerFaultDao struct {
	IBaseDao[do.ServerFaultDO, uint]
}

func NewServerFaultDao() *ServerFaultDao {
	return &ServerFaultDao{
		IBaseDao: &BaseDao[do.ServerFaultDO, uint]{
			Order: "start_time desc",
		},
	}
}

// 记录服务器故障
func (d ServerFaultDao) RecordFault(serverId uint, remark string) {
	d.RecordFaultAndTime(serverId, remark, time.Now())
}

func (d ServerFaultDao) RecordFaultAndTime(serverId uint, remark string, startTime time.Time) {
	var total int64
	d.DB().Model(&do.ServerFaultDO{}).Where("server_id = ? and end_time is null", serverId).Count(&total)
	if total > 0 {
		return
	}

	d.DB().Save(&do.ServerFaultDO{
		ServerId:  serverId,
		StartTime: time2.Time(startTime),
		Remark:    remark,
	})
}

// 服务器故障恢复
func (d ServerFaultDao) RecoverFault(serverId uint) {
	var serverFault do.ServerFaultDO
	d.DB().Where("server_id = ? and end_time is null", serverId).First(&serverFault)
	if serverFault.ID == 0 {
		return
	}

	endTimeV, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))
	endTime := time2.Time(endTimeV)

	fmt.Println("服务器故障恢复:", serverId, serverFault.StartTime)
	serverFault.EndTime = endTime
	serverFault.Duration = int64(time.Time(endTime).Sub(time.Time(serverFault.StartTime)).Seconds())
	d.DB().Save(&serverFault)
}
