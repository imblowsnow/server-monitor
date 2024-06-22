package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server-monitor/pkg/server/common/entity/bo"
	"server-monitor/pkg/server/dal/dao"
)

type ConfigController struct {
	configDao *dao.ConfigDao
}

func NewConfigController() *ConfigController {
	configDao := dao.NewConfigDao()
	configDao.InitConfig()
	return &ConfigController{configDao: configDao}
}
func (controller ConfigController) GetSetting(context *gin.Context) interface{} {
	// 判断是否登录
	context.Get("user")

	return controller.configDao.GetAllConfig()
}

func (controller ConfigController) UpdateSetting(context *gin.Context) interface{} {
	var configBo bo.ConfigBo
	err := context.ShouldBindJSON(&configBo)
	if err != nil {
		return fmt.Errorf("保存配置失败:%v", err)
	}
	controller.configDao.SaveConfig(configBo)

	return nil
}
