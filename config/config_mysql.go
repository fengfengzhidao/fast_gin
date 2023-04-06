package config

type Mysql struct {
	Host         string `yaml:"host"`                                 // 服务器地址:端口
	Port         string `yaml:"port"`                                 //端口
	Config       string `yaml:"config"`                               // 高级配置
	DB           string `yaml:"db"`                                   // 数据库名
	Username     string `yaml:"username"`                             // 数据库用户名
	Password     string `yaml:"password"`                             // 数据库密码
	MaxIdleConns int    `json:"max-idle-conns" yaml:"max-idle-conns"` // 空闲中的最大连接数
	MaxOpenConns int    `json:"max-open-conns" yaml:"max-open-conns"` // 打开到数据库的最大连接数
	LogMode      string `yaml:"log-mode"`                             // 是否开启Gorm全局日志
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.DB + "?" + m.Config
}

func (m *Mysql) GetLogMode() string {
	return m.LogMode
}
