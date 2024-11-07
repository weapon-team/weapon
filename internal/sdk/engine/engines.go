package engine

import (
	"time"

	"github.com/casbin/casbin/v2"
	"github.com/mojocn/base64Captcha"
	"github.com/redis/go-redis/v9"
	"xorm.io/xorm"

	"github.com/weapon-team/weapon/internal/sdk/captcha"
)

// Engines 依赖引擎封装
// desc: Engines将在本系统中传递,而不再单独传递某个依赖
// 如有其它引擎,加入变量即可
type Engines struct {
	orm     *xorm.Engine           // orm框架
	rdb     *redis.Client          // redis客户端
	enf     *casbin.SyncedEnforcer // casbin Enforcer
	captcha *base64Captcha.Captcha // 验证码处理器
}

// InitEngines 初始化依赖引擎
func InitEngines() (*Engines, error) {

	// 数据库引擎
	orm, err := initOrm()
	if err != nil {
		return nil, err
	}
	// redis客户端
	rdb, err := initRedisClient()
	if err != nil {
		return nil, err
	}
	// casbin Enforcer
	enf, err := initCasbinEnforcer(orm)
	if err != nil {
		return nil, err
	}
	// 验证码
	captchaEngine := captcha.NewDigitCaptcha(rdb, time.Minute*1)
	return &Engines{
		orm:     orm,
		rdb:     rdb,
		enf:     enf,
		captcha: captchaEngine,
	}, nil
}

func NewEngines(orm *xorm.Engine, rdb *redis.Client, enf *casbin.SyncedEnforcer) *Engines {
	return &Engines{
		orm: orm,
		rdb: rdb,
		enf: enf,
	}
}

// Orm Xorm
func (dep *Engines) Orm() *xorm.Engine {
	return dep.orm
}

// Redis client
func (dep *Engines) Redis() *redis.Client {
	return dep.rdb
}

// Casbin Enforcer
func (dep *Engines) Casbin() *casbin.SyncedEnforcer {
	return dep.enf
}

// Captcha 验证码
func (dep *Engines) Captcha() *base64Captcha.Captcha {
	return dep.captcha
}
