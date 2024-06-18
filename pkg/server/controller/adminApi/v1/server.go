package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"server-monitor/pkg/server/common/entity/vo"
	"server-monitor/pkg/server/common/enum"
	"server-monitor/pkg/server/controller/base"
	"server-monitor/pkg/server/dal/dao"
	"server-monitor/pkg/server/dal/do"
	"strconv"
)

type ServerController struct {
	serverDao *dao.ServerDao
	base.ICrudController[do.ServerDO, uint]
}

func NewServerController() *ServerController {
	var serverDao = dao.NewServerDao()
	return &ServerController{
		serverDao: serverDao,
		ICrudController: base.CrudController[do.ServerDO, uint]{
			Dao: serverDao,
		},
	}
}

func (s ServerController) GetServerInfo(context *gin.Context) interface{} {
	idParam := context.Param("id")
	id, _ := strconv.ParseUint(idParam, 10, 64)
	return s.serverDao.GetServerInfo(uint(id))
}

func (s ServerController) Create(context *gin.Context) interface{} {

	var newItem *vo.AddServerVO
	if err := context.ShouldBindJSON(&newItem); err != nil {
		return err
	}
	if newItem.GroupID == 0 {
		return fmt.Errorf("分组不能为空")
	}
	if newItem.Name == "" {
		return fmt.Errorf("服务器名称不能为空")
	}

	var serverDO = &do.ServerDO{
		GroupID:   newItem.GroupID,
		Key:       uuid.New().String(),
		Name:      newItem.Name,
		Remark:    newItem.Remark,
		ShowIndex: newItem.ShowIndex,
	}
	return s.GetDao().Add(serverDO)
}

func (s ServerController) GetServerStatisticsList(context *gin.Context) interface{} {
	idParam := context.Param("id")
	id, _ := strconv.ParseUint(idParam, 10, 64)
	typeParam, _ := strconv.Atoi(context.Query("type"))
	monitorDurationEnum := enum.FindMonitorDurationEnumByType(typeParam)
	return s.serverDao.GetMonitorServerStatisticsList(uint(id), monitorDurationEnum)
}
