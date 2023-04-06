package global

import (
	"fast_gin/config"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net"
)

type TcpServer struct {
	Conn   *net.TCPConn
	IP     string `json:"ip"`
	Port   int    `json:"port"`
	Status bool   `json:"status"` // 连接状态
}

var (
	DB     *gorm.DB
	CONFIG config.Server
	LOG    *logrus.Logger
)
