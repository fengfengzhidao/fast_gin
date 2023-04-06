package flags

import (
	"fast_gin/global"
)

func Makemigrations() {
	var err error
	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate()
	if err != nil {
		global.LOG.Error("[ error ] 生成数据库表结构失败")
		return
	}
	global.LOG.Info("[ success ] 生成数据库表结构成功！")
}
