package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
	"ums/utils"
)

var DB *gorm.DB
var err error

func init() {
	db, _ := gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			utils.DbUser,
			utils.DbPassword,
			utils.DbHost,
			utils.DbPort,
			utils.DbName),
		DefaultStringSize: 191,
	}), &gorm.Config{ // 自定义配置
		// 关闭默认的事务
		SkipDefaultTransaction: false,
		// 禁用外键约束，目前项目开发约束不通过数据库控制，而是通过代码控制，提供数据库性能
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix:   "t_", // 表名前缀
			SingularTable: true, // 使用单数表名
		},
	})

	// 常规数据库接口sql.DB, 配置连接池信息
	sqlDB, _ := db.DB()
	// 设置连接池中空闲连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// 设置打开数据库连接最大数量
	sqlDB.SetMaxOpenConns(100)
	// 设置连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Second * 10)

	DB = db
}
