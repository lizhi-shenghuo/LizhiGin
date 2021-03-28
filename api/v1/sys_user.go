package v1

import (
	"LizhiGin/global"
	"LizhiGin/model"
	"LizhiGin/model/request"
	"LizhiGin/model/response"
	"LizhiGin/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"LizhiGin/middleware"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"time"
)

func Login(c *gin.Context) {
	var L request.Login
	err := c.ShouldBindJSON(&L)
	if err != nil {
		return
	}
	if err := utils.Verify(L, utils.LoginVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
}


// 登录以后签发jwt
func tokenNext(c *gin.Context, user model.SysUser) {
	j := &middleware.JWT{SigningKey: []byte(global.LizhiConfig.JWT.SigningKey)}
	claims := request.CustomClaims{
		UUID:           user.UUID,
		ID:             user.ID,
		Username:       user.Username,
		NickName:       user.NickName,
		AuthorityId:    "",
		BufferTime:     0,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() - 1000,
			NotBefore: time.Now().Unix() + global.LizhiConfig.JWT.ExpiresTime,
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		global.LizhiLog.Error("获取token失败",zap.Any("err", err))
		response.FailWithMessage("获取token失败",c)
		return
	}
	if !global.LizhiConfig.
}