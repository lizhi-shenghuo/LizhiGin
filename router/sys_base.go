package router

import (
	"LizhiGin/api/v1"
	"LizhiGin/middleware"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(r *gin.RouterGroup) *gin.IRouter {
	BaseRouter := r.Group("base").Use(middleware.)
}