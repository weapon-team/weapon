package helper

// LoginParam 账号登录请求参数
type LoginParam struct {
	Username  string `form:"username" json:"username" validate:"required" error:"请输入用户名"`
	Password  string `form:"password" json:"password" validate:"required" error:"请输入密码"`
	CaptchaId string `json:"captchaId" validate:"required" error:"无效的验证码ID"`
	Captcha   string `form:"captcha" json:"captcha" validate:"required" error:"请输入验证码"`
}

// CreateUserParam 创建用户请求参数
type CreateUserParam struct {
	Username string `form:"username" json:"username" validate:"required" error:"请输入用户名"`
	Nickname string `form:"nickname" json:"nickname" validate:"required" error:"请输入昵称"`
	Password string `form:"password" json:"password" validate:"required" error:"请输入密码"`
	Gender   uint   `form:"gender" json:"gender" validate:"oneof=0 1 2" error:"请选择性别"`
}

// UpdateUserParam 修改用户请求参数
type UpdateUserParam struct {
	Id       int64  `form:"id" json:"id" validate:"required" error:"无效的用户id"`
	Username string `form:"username" json:"username" validate:"required" error:"请输入用户名"`
	Nickname string `form:"nickname" json:"nickname" validate:"required" error:"请输入昵称"`
	Password string `form:"password" json:"password" validate:"required" error:"请输入密码"`
	Gender   uint   `form:"gender" json:"gender" validate:"oneof=0 1 2" error:"请选择性别"`
}
