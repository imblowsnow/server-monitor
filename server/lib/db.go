package lib

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"server/config"
	"server/entity/do"
)

func init() {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&do.Server{}, &do.ServerState{}, &do.ServerInfo{}, &do.ServerFault{})

	config.Db = db

	// 初始化数据库
	fmt.Println("数据库初始化成功")
}
