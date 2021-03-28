package middleware

import (
	"LizhiGin/global"
	"LizhiGin/model/response"
	"github.com/gin-gonic/gin"
)


// 处理跨域请求,支持options访问
func NeedInit() gin.HandlerFunc {
	return func(c *gin.Context) {
		if global.LizhiDB == nil {
			response.OkWithDetail(gin.H{
				"needInit": true,
			}, "前往初始化数据库",c)
		} else {
			c.Next()
		}
	}
}
