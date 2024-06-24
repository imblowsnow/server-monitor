package v1

import (
	"github.com/gin-gonic/gin"
	"server-monitor/pkg/server/common/entity/bo"
	"server-monitor/pkg/server/common/enum"
	"server-monitor/pkg/server/common/utils"
	"server-monitor/pkg/server/dal/dao"
	"strconv"
)

type ServerController struct {
	serverDao *dao.ServerDao
}

func NewServerController() *ServerController {
	return &ServerController{
		serverDao: dao.NewServerDao(),
	}
}

func (s ServerController) GetServerInfo(context *gin.Context) interface{} {
	idParam := context.Param("id")
	id, _ := strconv.ParseUint(idParam, 10, 64)
	serverInfoBo := s.serverDao.GetServerInfo(uint(id))

	// 转换为 frontendServerInfoBO
	frontendServerInfoBO := &bo.FrontendServerInfoBO{}

	utils.CopyObject(serverInfoBo, frontendServerInfoBO)

	return frontendServerInfoBO
}

func (s ServerController) GetServerStatisticsList(context *gin.Context) interface{} {
	idParam := context.Param("id")
	id, _ := strconv.ParseUint(idParam, 10, 64)
	typeParam, _ := strconv.Atoi(context.Query("type"))
	monitorDurationEnum := enum.FindMonitorDurationEnumByType(typeParam)
	return s.serverDao.GetMonitorServerStatisticsList(uint(id), monitorDurationEnum)
}
