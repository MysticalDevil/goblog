package model

import (
	"goblog/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	var err error
	dsn := "host=localhost user=postgres password=112233 dbname=goblog port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	config := postgres.New(postgres.Config{
		DSN: dsn,
	})

	DB, err = gorm.Open(config, &gorm.Config{})
	logger.LogError(err)

	return DB
}
