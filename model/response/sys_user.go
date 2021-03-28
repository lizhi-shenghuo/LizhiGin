package response

import "LizhiGin/model"

type SysUserResponse struct {
	User model.SysUser `json:"user"`
 }

type LoginResponse struct {
	User model.SysUser `json:"user"`
	Token string `json:"token"`
	ExpireAt int64 `json:"expire_at"`
}