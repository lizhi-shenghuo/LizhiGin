package global

import (
	"LizhiGin/config"
	"go.uber.org/zap"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	LizhiDB *gorm.DB
	LizhiRedis *redis.Client
	LizhiConfig config.Server
	LizhiViper *viper.Viper
	LizhiLog *zap.Logger
)