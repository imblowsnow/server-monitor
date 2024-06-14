package config

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"server-monitor/pkg/server/constants"
)

func init() {
	fmt.Println("数据库初始化开始")

	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("failed to connect database")
	}

	constants.DB = db

	// 初始化数据库
	fmt.Println("数据库初始化成功")
}
