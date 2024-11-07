package captcha

import (
	"testing"
	"time"

	"github.com/mojocn/base64Captcha"
)

func TestCaptcha(t *testing.T) {
	driver := base64Captcha.NewDriverDigit(100, 200, 6, 0.6, 8)
	store := base64Captcha.NewMemoryStore(100, time.Second*6)
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, ans, err := c.Generate()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("id: %v, b64s: %v, ans: %v", id, b64s, ans)
}
