package engine

import (
	"github.com/casbin/casbin/v2"
	"github.com/redis/go-redis/v9"
	"xorm.io/xorm"
)

// Engines 依赖引擎封装
// desc: Engines将在本系统中传递,而不再单独传递某个依赖
// 如有其它引擎,加入变量即可
type Engines struct {
	orm *xorm.Engine           // orm框架
	rdb *redis.Client          // redis客户端
	enf *casbin.SyncedEnforcer // casbin Enforcer
}

// InitEngines 初始化依赖引擎
func InitEngines() (*Engines, error) {

	orm, err := initOrm()
	if err != nil {
		return nil, err
	}
	rdb, err := initRedisClient()
	if err != nil {
		return nil, err
	}
	enf, err := initCasbinEnforcer(orm)
	if err != nil {
		return nil, err
	}
	return &Engines{
		orm: orm,
		rdb: rdb,
		enf: enf,
	}, nil
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
