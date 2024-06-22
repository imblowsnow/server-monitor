package dao

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"server-monitor/pkg/server/common/entity/bo"
	"server-monitor/pkg/server/config"
	"server-monitor/pkg/server/dal/do"
)

type ConfigDao struct {
	IBaseDao[do.ConfigDO, int]
}

func NewConfigDao() *ConfigDao {
	return &ConfigDao{
		IBaseDao: &BaseDao[do.ConfigDO, int]{},
	}
}

// 获取所有配置
func (dao ConfigDao) GetAllConfig() bo.ConfigBo {
	var list []do.ConfigDO
	dao.DB().Find(&list)
	result := make(map[string]string)
	for _, item := range list {
		result[item.Name] = item.Value
	}
	// map 转 bo
	byte, _ := json.Marshal(result)
	var bo bo.ConfigBo
	json.Unmarshal(byte, &bo)
	return bo
}

// 保存配置
func (dao ConfigDao) SaveConfig(configBo bo.ConfigBo) error {
	// configBo 转 map
	configByte, err := json.Marshal(configBo)
	if err != nil {
		return err
	}
	var configMap map[string]string
	err = json.Unmarshal(configByte, &configMap)
	if err != nil {
		return err
	}
	for key, value := range configMap {
		var configDo do.ConfigDO
		dao.DB().Where("name = ?", key).First(&configDo)
		if configDo.ID > 0 {
			configDo.Value = value
			result := dao.DB().Save(&configDo)
			if result.Error != nil {
				return result.Error
			}
		}
	}

	config.Config = dao.GetAllConfig()
	return nil
}

func (dao ConfigDao) initConfigData(configBo bo.ConfigBo) error {
	// configBo 转 map
	configByte, err := json.Marshal(configBo)
	if err != nil {
		return err
	}
	var configMap map[string]string
	err = json.Unmarshal(configByte, &configMap)
	if err != nil {
		return err
	}
	for key, value := range configMap {
		var configDo do.ConfigDO
		dao.DB().Where("name = ?", key).First(&configDo)
		if configDo.ID > 0 {
		} else {
			result := dao.DB().Create(&do.ConfigDO{Name: key, Value: value})
			if result.Error != nil {
				return result.Error
			}
		}
	}

	config.Config = dao.GetAllConfig()
	return nil
}

func (dao ConfigDao) InitConfig() {
	configBo := bo.ConfigBo{
		Secret:           uuid.New().String(),
		SiteName:         "Server Monitor",
		FrontendTemplate: "default",
		BackendTemplate:  "default",
		AdminUsername:    "admin",
		AdminPassword:    "admin",
	}
	err := dao.initConfigData(configBo)
	if err != nil {
		fmt.Println("初始化默认配置失败")
		panic(err)
	}
}
