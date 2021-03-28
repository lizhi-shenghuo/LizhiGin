package initialize


import (

	"LizhiGin/global"
	"LizhiGin/middleware"
	"LizhiGin/router"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// 初始化总路由

func Routers() *gin.Engine {
	var Router = gin.Default()
	Router.StaticFS(global.LizhiConfig.Local.Path, http.Dir(global.LizhiConfig.Local.Path))
	global.LizhiLog.Info("use middleware logger")

	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	global.LizhiLog.Info("register swagger handler")

	global.LizhiLog.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用
	PublicGroup := Router.Group("")
	{
		router.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
		router.InitInitRouter(PublicGroup) // 自动初始化相关
	}
}