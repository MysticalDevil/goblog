package bootstrap

import (
	"goblog/pkg/model"
	"time"
)

func SetupDB() {
	// 建立数据库连接池
	db := model.ConnectDB()

	// 命令行打印数据库请求信息
	sqlDB, _ := db.DB()

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(time.Minute * 5)
}
