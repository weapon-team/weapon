package service

import "github.com/weapon-team/weapon/internal/sdk/engine"

// CaptchaService 验证码服务
type CaptchaService struct {
	*engine.Engines
}

func NewCaptchaService(deps *engine.Engines) *CaptchaService {
	return &CaptchaService{deps}
}

// Generate 生成验证码
func (c *CaptchaService) Generate() (id string, b64s string, err error) {
	id, b64s, _, err = c.Captcha().Generate()
	return
}

// Verify 验证验证码
func (c *CaptchaService) Verify(id, answer string) bool {
	return c.Captcha().Verify(id, answer, true)
}
