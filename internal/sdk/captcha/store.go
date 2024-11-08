package captcha

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/mojocn/base64Captcha"
	"github.com/redis/go-redis/v9"
)

// Store 验证码存储器
type Store struct {
	rc         *redis.Client
	expiration time.Duration
}

var _ base64Captcha.Store = (*Store)(nil)

func NewStore(rc *redis.Client, expiration time.Duration) *Store {
	return &Store{rc: rc, expiration: expiration}
}

func (c *Store) Set(id string, value string) error {
	if id == "" || value == "" {
		return errors.New("captcha: id or value is empty")
	}
	return c.rc.Set(context.Background(), fmt.Sprintf("captcha_%s", id), value, c.expiration).Err()
}

func (c *Store) Get(id string, clear bool) string {
	captchaKey := fmt.Sprintf("captcha_%s", id)
	val := c.rc.Get(context.Background(), captchaKey).String()
	if clear && val != "" {
		c.rc.Del(context.Background(), captchaKey)
	}
	return val
}

func (c *Store) Verify(id, answer string, clear bool) bool {
	if id == "" || answer == "" {
		return false
	}
	v := c.Get(id, clear)
	return v != "" && v == answer
}
