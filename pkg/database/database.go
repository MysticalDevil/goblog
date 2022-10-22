package database

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"goblog/pkg/logger"
	"time"
)

var DB *sql.DB

func Initialize() {
	initDB()
	createTables()
}

func initDB() {
	var err error
	config := mysql.Config{
		User:                 "root",
		Passwd:               "112233",
		Addr:                 "127.0.0.1:3308",
		Net:                  "tcp",
		DBName:               "goblog",
		AllowNativePasswords: true,
	}

	// 准备数据库连接池
	db, err = sql.Open("mysql", config.FormatDSN())
	logger.LogError(err)

	// 设置最大连接数
	db.SetMaxOpenConns(25)
	// 设置最大空闲连接数
	db.SetMaxIdleConns(25)
	// 设置每个连接的过期时间
	db.SetConnMaxIdleTime(time.Minute * 5)

	// 尝试连接
	err = db.Ping()
	logger.LogError(err)
}

func createTables() {
	createArticlesSQL := `CREATE TABLE IF NOT EXISTS articles(
    id BIGINT(20) PRIMARY KEY AUTO_INCREMENT NOT NULL,
    title VARCHAR(255) COLLATE utf8mb4_unicode_ci NOT NULL,
    body LONGTEXT COLLATE utf8mb4_unicode_ci
    );`

	_, err := db.Exec(createArticlesSQL)
	logger.LogError(err)
}
