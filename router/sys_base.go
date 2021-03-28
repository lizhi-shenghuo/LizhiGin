package router

import (
	"LizhiGin/api/v1"
	"LizhiGin/middleware"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(r *gin.RouterGroup) gin.IRoutes {
	BaseRouter := r.Group("base").Use(middleware.NeedInit())
	{
		BaseRouter.POST("login", v1.Login)
	}
	return BaseRouter
}

