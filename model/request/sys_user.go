package request

import uuid "github.com/satori/go.uuid"

// User Register structure
type Register struct {
	Username string `json:"username"`
	Password string `json:"password"`
	NickName string `json:"nick_name"`
	HeaderImg string `json:"header_img" gorm:"default:'http://www.henrongyi.top/avatar/lufu.jpg'"`
	AuthorityId string `json:"authorityId" gorm:"default:888"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Captcha string `json:"captcha"`
	CaptachaId string `json:"captacha_id"`
}

// Modify password structure
type ChangePasswordStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
	NewPassword string `json:"new_password"`
}


// Modify user's auth structure
type SetUserAuth struct {
	UUID uuid.UUID `json:"uuid"`
	AuthorityId string `json:"authority_id"`
}