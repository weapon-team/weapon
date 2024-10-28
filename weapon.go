package main

import (
	"context"
	"fmt"

	"github.com/casbin/casbin/v2"
	xormadapter "github.com/casbin/xorm-adapter/v2"
	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
	"xorm.io/xorm"

	"github.com/weapon-team/weapon/internal"
	"github.com/weapon-team/weapon/internal/sdk/runtime"
	"github.com/weapon-team/weapon/internal/sdk/xormlog"
	"github.com/weapon-team/weapon/pkg/banner"
	"github.com/weapon-team/weapon/pkg/logs"
)

func main() {

	fmt.Println("\nWelcome to :")
	fmt.Println(banner.WEAPON)

	// 1. 初始化配置
	if err := runtime.InitConfig("config/config.yml"); err != nil {
		panic(err)
	}

	fmt.Println("Weapon is starting...")
	fmt.Println("Config: ", runtime.Setting.String())

	// 2. 初始化全局log
	logCfg := runtime.Setting.Log
	logs.Init(logCfg.Level, logCfg.Path, logCfg.Interval)

	// 3. 数据库引擎
	orm, err := xorm.NewEngine(runtime.Setting.DataSource.Driver, runtime.Setting.DataSource.DSN)
	if err != nil {
		panic(err.Error())
	}
	orm.SetLogger(&xormlog.XormLogger{})
	orm.ShowSQL(logCfg.ShowSQL)
	if err := orm.Ping(); err != nil {
		panic(err)
	}

	// 4. redis缓存
	rdb := redis.NewClient(&redis.Options{
		Addr:           runtime.Setting.Redis.Addr,
		Username:       runtime.Setting.Redis.Username,
		Password:       runtime.Setting.Redis.Password,
		MaxActiveConns: runtime.Setting.Redis.MaxActive,
	})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		panic(fmt.Errorf("failed to connect to redis: %w", err))
	}

	// 5. casbin
	engine, err := xormadapter.NewAdapterByEngine(orm)
	if err != nil {
		panic(err)
	}
	enf, err := casbin.NewSyncedEnforcer("config/rbac_models.conf", engine)
	if err != nil {
		panic(err)
	}
	err = enf.LoadPolicy()
	if err != nil {
		panic(err)
	}

	// 6. 启动路由
	internal.StartRouter(orm, rdb, enf)
}
