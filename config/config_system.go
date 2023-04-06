package config

import "fmt"

type System struct {
	Env     string `yaml:"env"`     // 环境值 区分开发环境测试环境
	IP      string `yaml:"ip"`      // 项目运行的ip
	Port    int    `yaml:"port"`    // 项目运行的端口
	Version string `yaml:"version"` // 项目的版本
}

func (s System) GetAddr() string {
	return fmt.Sprintf("%s:%d", s.IP, s.Port)
}

func (s System) GetHttpAddr() string {
	return fmt.Sprintf("http://%s", s.GetAddr())
}
