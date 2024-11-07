package captcha

import (
	"time"

	"github.com/mojocn/base64Captcha"
	"github.com/redis/go-redis/v9"
)

// NewDigitCaptcha 数字验证码
func NewDigitCaptcha(rc *redis.Client, expiration time.Duration) *base64Captcha.Captcha {
	store := NewStore(rc, expiration)
	driver := base64Captcha.NewDriverDigit(100, 200, 6, 0.6, 15)
	return base64Captcha.NewCaptcha(driver, store)
}
