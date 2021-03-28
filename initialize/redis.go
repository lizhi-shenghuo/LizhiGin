package initialize

import (
	"LizhiGin/global"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

func Redis()  {
	redisCfg := global.LizhiConfig.Redis
	client := redis.NewClient(&redis.Options{
		Addr:               redisCfg.Addr,
		Password:           redisCfg.Password,
		DB:                 redisCfg.DB,
	})
	pong, err := client.Ping().Result()
	if err != nil {
		global.LizhiLog.Error("redis connect ping failed, err:", zap.Any("err", err))
	} else {
		global.LizhiLog.Info("redis connect ping response", zap.String("pong", pong))
		global.LizhiRedis = client
	}
}