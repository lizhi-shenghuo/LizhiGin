package service

import (
	"errors"
	"LizhiGin/global"
	"LizhiGin/model"
	"gorm.io/gorm"
	"time"
)


// 拉黑jwt
func JsonInBlacklist(jwtList model.JwtBlacklist) (err error) {
	err = global.LizhiDB.Create(&jwtList).Error
	return
}

//@description: 判断JWT是否在黑名单内部
func IsBlacklist(jwt string) bool {
	isNotFound := errors.Is(global.LizhiDB.Where("jwt = ?", jwt).First(&model.JwtBlacklist{}).Error, gorm.ErrRecordNotFound)
	return !isNotFound
}

//@description: 从redis取jwt
func GetRedisJWT(userName string) (redisJWT string, err error)  {
	redisJWT, err  = global.LizhiRedis.Get(userName).Result()
	return redisJWT, err
}

//@description: jwt存入redis并设置过期时间
func SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	timer := time.Duration(global.LizhiConfig.JWT.ExpiresTime) * time.Second
	err = global.LizhiRedis.Set(userName, jwt, timer).Err()
	return err
}