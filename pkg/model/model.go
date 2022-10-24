package model

import (
	"fmt"
	c "goblog/pkg/config"
	"goblog/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	var err error
	//dsn := "host=localhost user=postgres password=112233 dbname=goblog port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable TimeZone=%v",
		c.GetString("database.postgresql.host"),
		c.GetString("database.postgresql.port"),
		c.GetString("database.postgresql.username"),
		c.GetString("database.postgresql.password"),
		c.GetString("database.postgresql.database"),
		c.GetString("database.postgresql.timezone"),
	)

	var level gormLogger.LogLevel
	if c.GetBool("app.debug") {
		level = gormLogger.Warn
	} else {
		level = gormLogger.Error
	}

	config := postgres.New(postgres.Config{
		DSN: dsn,
	})

	DB, err = gorm.Open(config, &gorm.Config{
		Logger: gormLogger.Default.LogMode(level),
	})
	logger.LogError(err)

	return DB
}
