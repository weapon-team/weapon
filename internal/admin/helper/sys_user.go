package helper

// AccountLoginParam 账号登录请求参数
type AccountLoginParam struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Captcha  string `form:"captcha" json:"captcha" binding:"required"`
	Uuid     string `form:"uuid" json:"uuid" binding:"required"`
}
