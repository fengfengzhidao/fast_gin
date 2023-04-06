package config

type Server struct {
	Mysql   Mysql   `yaml:"mysql"`
	Loggers Loggers `yaml:"logger"`
	System  System  `yaml:"system"`
	Jwy     Jwy     `yaml:"jwy"`
}
