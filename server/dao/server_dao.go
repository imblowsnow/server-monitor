package dao

import (
	"server/config"
	"server/entity/do"
)

type ServerDao struct {
}

func GetServerDao() *ServerDao {
	return &ServerDao{}
}

// 保存服务器信息
func (s *ServerDao) AddServer(server do.Server) error {
	result := config.Db.Create(&server)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// 通过key服务器获取信息
func (s *ServerDao) GetServerByKey(key string) (do.Server, error) {
	server := do.Server{}
	result := config.Db.Where("key = ?", key).Limit(1).Find(&server)
	if result.Error != nil {
		return do.Server{}, result.Error
	}
	return server, nil
}

// 通过id获取服务器信息
func (s *ServerDao) GetServerById(id uint) (do.Server, error) {
	server := do.Server{}
	result := config.Db.Where("id = ?", id).Limit(1).Find(&server)
	if result.Error != nil {
		return do.Server{}, result.Error
	}
	return server, nil
}

// 删除服务器
func (s *ServerDao) DeleteServer(id uint) error {
	result := config.Db.Where("id = ?", id).Delete(do.Server{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// 更新服务器信息
func (s *ServerDao) UpdateServer(server do.Server) {
	config.Db.Save(&server)
}

// 通过排序获取服务器列表
func (s *ServerDao) GetServerList(hideIndex bool) []do.Server {
	var list []do.Server
	whers := do.Server{}
	if hideIndex {
		whers.Hide = 0
	}
	config.Db.Where(whers).Order("`index` desc").Find(&list)
	return list
}
