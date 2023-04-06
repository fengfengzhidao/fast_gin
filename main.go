package main

import (
	"fast_gin/core"
	"fast_gin/docs"
	"fast_gin/flags"
	"fast_gin/global"
	"fast_gin/routers"
	"net/http"
	"time"
)

// @title API文档
// @version 1.0
// @description API文档
// @host 127.0.0.1:8080
// @BasePath /
func main() {
	global.CONFIG = core.ConfInit()
	global.LOG = core.InitLogger()
	global.DB = core.Gorm()

	//获取命令参数
	option := flags.Parse()
	if flags.SwitchOption(option) {
		return
	}

	route := routers.InitRouter()

	addr := global.CONFIG.System.GetAddr()
	docs.SwaggerInfo.Host = addr
	s := &http.Server{
		Addr:           addr,
		Handler:        route,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	global.LOG.Infof("服务运行在：http://%s", addr)
	global.LOG.Infof("api接口文档运行在：http://%s/swagger/index.html", addr)

	s.ListenAndServe().Error()
}
