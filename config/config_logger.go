package config

type Loggers struct {
	Level        string `yaml:"level"`          // 级别
	Prefix       string `yaml:"prefix"`         // 日志前缀
	Director     string `yaml:"director"`       // 日志文件夹
	ShowLine     bool   `yaml:"show-line"`      // 显示行
	LogInConsole bool   `yaml:"log-in-console"` // 输出控制台
}
