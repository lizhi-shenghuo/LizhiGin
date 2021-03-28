package utils

var (
	IdVerify  = Rules{"ID": {NotEmpty()}}
	LoginVerify  = Rules{"CaptchaId": {NotEmpty()}, "Captcha": {NotEmpty()}, "Username": {NotEmpty()}, "Password": {NotEmpty()}}
)
