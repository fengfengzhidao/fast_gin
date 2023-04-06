package routers

import (
	"fast_gin/global"
	res "fast_gin/models/common"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

type RouterGroup struct {
	*gin.RouterGroup
}

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	gin.SetMode(global.CONFIG.System.Env)
	router := gin.Default()
	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	PublicGroup := router.Group("api")
	PublicGroup.GET("/", func(c *gin.Context) {
		res.OkWithMessage("成功", c)
	})

	apiRouterGroup := router.Group("api")
	routerGroupApp := RouterGroup{apiRouterGroup}
	// 系统配置api
	routerGroupApp.UserRouter()

	return router
}
