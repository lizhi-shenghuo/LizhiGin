package model

import "LizhiGin/global"

type JwtBlacklist struct {
	global.LizhiModel
	Jwt string `gorm:"type:text";comment:jwt`
}
