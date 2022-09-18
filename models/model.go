package models

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// 数据库连接
// 定义数据地址
var Dsn string = "root:qiudaoyu@tcp(127.0.0.1:3306)/qiudaoyu?charset=utf8mb4&parseTime=True&loc=Local"
var Conn *gorm.DB = nil

func InitDb() {
	var err error
	Conn, err = gorm.Open(mysql.Open(Dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	sqlDB, _ := Conn.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(200)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)
}
