package v1

import (
	"LizhiGin/global"
	"LizhiGin/middleware"
	"LizhiGin/model"
	"LizhiGin/model/request"
	"LizhiGin/model/response"
	"LizhiGin/service"
	"LizhiGin/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
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
	j := &middleware.JWT{SigningKey: []byte(global.LizhiConfig.JWT.SigningKey)}  // 唯一签名
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
			Issuer: "sunzhou",
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		global.LizhiLog.Error("获取token失败",zap.Any("err", err))
		response.FailWithMessage("获取token失败",c)
		return
	}
	if !global.LizhiConfig.System.UseMultipoint {
		response.OkWithDetail(response.LoginResponse{
			User: user,
			Token: token,
			ExpireAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功",c)
		return
	}
	if jwtStr, err := service.GetRedisJWT(user.Username); err == redis.Nil {
		if err := service.SetRedisJWT(token, user.Username); err != nil {
			global.LizhiLog.Error("设置登录状态失败", zap.Any("err", err))
			response.FailWithMessage("设置登录状态失败",c)
			return
		}
		response.OkWithDetail(response.LoginResponse{
			User:     user,
			Token:    token,
			ExpireAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	} else if err != nil {
		global.LizhiLog.Error("设置登录状态失败", zap.Any("err", err))
		response.FailWithMessage("设置登录状态失败", c)
	} else {
		var blackJWT model.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := service.JsonInBlacklist(blackJWT); err != nil {
			response.FailWithMessage("jwt拉黑失败", c)
			return
		}
		if err := service.SetRedisJWT(token, user.Username); err != nil {
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetail(response.LoginResponse{
			User:     user,
			Token:    token,
			ExpireAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	}
}