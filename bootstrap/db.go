package bootstrap

import (
	"goblog/app/models/article"
	"goblog/app/models/category"
	"goblog/app/models/user"
	"goblog/pkg/config"
	"goblog/pkg/logger"
	"goblog/pkg/model"
	"gorm.io/gorm"
	"time"
)

func SetupDB() {
	// 建立数据库连接池
	db := model.ConnectDB()

	// 命令行打印数据库请求信息
	sqlDB, _ := db.DB()

	sqlDB.SetMaxOpenConns(config.GetInt("database.postgresql.max_open_connections"))
	sqlDB.SetMaxIdleConns(config.GetInt("database.postgresql.max_idle_connections"))
	sqlDB.SetConnMaxLifetime(time.Duration(config.GetInt("database.postgresql.max_life_seconds"))*time.Second)

	migration(db)
}

func migration(db *gorm.DB) {
	err := db.AutoMigrate(
		&user.User{},
		&article.Article{},
		&category.Category{},
	)
	logger.LogError(err)
}
