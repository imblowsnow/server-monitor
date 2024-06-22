package v1

import (
	"github.com/gin-gonic/gin"
	"server-monitor/pkg/server/common/entity/bo"
	"server-monitor/pkg/server/dal/dao"
)

type ConfigController struct {
	configDao *dao.ConfigDao
}

func NewConfigController() *ConfigController {
	return &ConfigController{
		configDao: dao.NewConfigDao(),
	}
}

func (controller ConfigController) GetConfig(context *gin.Context) interface{} {

	configBo := controller.configDao.GetAllConfig()

	return bo.FrontConfigBo{
		SiteName: configBo.SiteName,
	}
}
