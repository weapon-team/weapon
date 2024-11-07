package helper

// LoginParam 账号登录请求参数
type LoginParam struct {
	Username string `form:"username" json:"username" validate:"required" error:"invalid username"`
	Password string `form:"password" json:"password" validate:"required" error:"invalid password"`
	Captcha  string `form:"captcha" json:"captcha" validate:"required" error:"invalid captcha"`
	Uuid     string `form:"uuid" json:"uuid" validate:"required" error:"invalid uuid"`
}
