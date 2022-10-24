package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"goblog/pkg/logger"
	"time"
)

var DB *sql.DB

func Initialize() {
	initDB()
}

func initDB() {
	var err error
	connStr := "host=localhost user=postgres password=112233 dbname=goblog port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	// 准备数据库连接池
	DB, err = sql.Open("postgres", connStr)
	logger.LogError(err)

	// 设置最大连接数
	DB.SetMaxOpenConns(25)
	// 设置最大空闲连接数
	DB.SetMaxIdleConns(25)
	// 设置每个连接的过期时间
	DB.SetConnMaxIdleTime(time.Minute * 5)

	// 尝试连接
	err = DB.Ping()
	logger.LogError(err)
}
