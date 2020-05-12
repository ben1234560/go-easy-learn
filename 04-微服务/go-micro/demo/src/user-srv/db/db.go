package db

import (
	_ "github.com/go-sql-driver/mysql"
	"jmoiron/sqlx"
)

// 声明数据库实例
var (
	db *sqlx.DB
)

func Init(mysqlDNS string) {
	// 获取连接
	db = sqlx.MustConnect("mysql", mysqlDNS)
	// 设置闲置的连接数
	db.SetMaxIdleConns(1)
	// 设置最大打开的连接数
	db.SetMaxOpenConns(3)
}
