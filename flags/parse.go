package flags

import (
	"fast_gin/global"
	"flag"
)

type Option struct {
	Version bool
	DB      bool
	Port    int
}

// Parse 解析命令行参数
func Parse() Option {
	version := flag.Bool("v", false, "项目版本")
	db := flag.Bool("db", false, "初始化数据库")
	port := flag.Int("p", 0, "端口")
	// 解析命令行参数写入注册的flag里
	flag.Parse()
	return Option{
		Version: *version,
		DB:      *db,
		Port:    *port,
	}
}

// SwitchOption 根据命令执行不同的函数
func SwitchOption(option Option) bool {
	if option.DB {
		Makemigrations()
		return true
	}
	if option.Version {
		Version()
		return true
	}
	if option.Port != 0 {
		global.CONFIG.System.Port = option.Port
		return false
	}
	return false
}
