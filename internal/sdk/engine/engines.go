package engine

import (
	"github.com/redis/go-redis/v9"
	"xorm.io/xorm"
)

// Engines 依赖封装
// desc: 如有其它引擎,加入变量即可
type Engines struct {
	orm *xorm.Engine
	rdb *redis.Client
}

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
