package helper

// LoginParam 账号登录请求参数
type LoginParam struct {
	Username  string `form:"username" json:"username" validate:"required" error:"invalid username"`
	Password  string `form:"password" json:"password" validate:"required" error:"invalid password"`
	CaptchaId string `json:"captchaId" validate:"required" error:"invalid captchaId"`
	Captcha   string `form:"captcha" json:"captcha" validate:"required" error:"invalid captcha"`
}
