package dep

import (
	"github.com/redis/go-redis/v9"
	"xorm.io/xorm"
)

// Dependency 依赖封装
type Dependency struct {
	orm *xorm.Engine
	rdb *redis.Client
}

func NewDependency(engine *xorm.Engine, rdb *redis.Client) *Dependency {
	return &Dependency{
		orm: engine,
		rdb: rdb,
	}
}

func (dep *Dependency) Orm() *xorm.Engine {
	return dep.orm
}

func (dep *Dependency) Redis() *redis.Client {
	return dep.rdb
}
