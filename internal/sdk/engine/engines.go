package engine

import (
	"github.com/redis/go-redis/v9"
	"xorm.io/xorm"
)

// Engines 依赖引擎封装
// desc: Engines将在本系统中传递,而不再单独传递某个依赖
// 如有其它引擎,加入变量即可
type Engines struct {
	orm *xorm.Engine  // orm框架
	rdb *redis.Client // redis客户端
}

// NewEngines 新建
func NewEngines(engine *xorm.Engine, rdb *redis.Client) *Engines {
	return &Engines{
		orm: engine,
		rdb: rdb,
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
