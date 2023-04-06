package core

import (
	"fast_gin/global"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

// Gorm 初始化数据库并产生数据库全局变量
// Author SliverHorn
func Gorm() *gorm.DB {
	return MysqlConnect()
}

func MysqlConnect() *gorm.DB {

	if global.CONFIG.Mysql.Host == "" {
		global.LOG.Warn("未配置mysql，取消gorm连接")
		return nil
	}
	dsn := global.CONFIG.Mysql.Dsn()

	var mysqlLogger logger.Interface
	if global.CONFIG.System.Env == "dev" {
		// 开发环境显示所有的sql
		mysqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		mysqlLogger = logger.Default.LogMode(logger.Error) // 只打印错误的sql
	}
	//global.MysqlLog = logger.Default.LogMode(logger.Info)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		global.LOG.Error(fmt.Sprintf("[%s] mysql连接失败", dsn))
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)               // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)              // 最多可容纳
	sqlDB.SetConnMaxLifetime(time.Hour * 4) // 连接最大复用时间，不能超过mysql的wait_timeout
	return db
}
