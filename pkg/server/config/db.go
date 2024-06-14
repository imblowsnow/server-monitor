package config

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"server-monitor/pkg/server/dal/do"
	"strings"
	"sync"
)

var _db *gorm.DB
var mutex sync.Mutex

func DB() *gorm.DB {
	if _db == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if _db == nil {
			initDb()
		}
		return _db
	}
	return _db
}

func init() {
	initDb()
}
func initDb() {
	fmt.Println("数据库初始化开始")

	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			NameReplacer:  strings.NewReplacer("DO", ""),
		},
	})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&do.ServerDO{}, &do.ServerInfoDO{}, &do.ServerStateDO{}, &do.ServerInfoDO{})
	db.AutoMigrate(&do.NotifyGroupDO{}, &do.NotifyChannelDO{}, &do.NotifyGroupChannelDO{}, &do.NotifyLogDO{})

	_db = db

	// 初始化数据库
	fmt.Println("数据库初始化成功")
}
