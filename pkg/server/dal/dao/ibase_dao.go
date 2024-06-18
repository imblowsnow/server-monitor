package dao

import (
	"gorm.io/gorm"
	"server-monitor/pkg/server/common/entity"
)

type IBaseDao[DO any, ID comparable] interface {
	GetList(querys map[string]interface{}) []DO
	Page(page *entity.Page[DO], querys map[string]interface{}) error
	Delete(id ID) error
	GetById(id ID) *DO
	Add(entity *DO) error
	UpdateById(id ID, entity *DO) error
	Update(entity *DO) error
	DB() *gorm.DB
}
