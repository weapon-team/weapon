package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"xorm.io/xorm"

	_ "github.com/go-sql-driver/mysql"

	"github.com/weapon-team/weapon/internal"
	"github.com/weapon-team/weapon/internal/sdk/runtime"
	"github.com/weapon-team/weapon/pkg/banner"
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

	// 2. 数据库引擎
	orm, err := xorm.NewEngine(runtime.Setting.DataSource.Driver, runtime.Setting.DataSource.DNS)
	if err != nil {
		panic(err.Error())
	}
	orm.ShowSQL(true)
	if err := orm.Ping(); err != nil {
		panic(err)
	}

	// 3. redis缓存
	rdb := redis.NewClient(&redis.Options{
		Addr:           runtime.Setting.Redis.Addr,
		Username:       runtime.Setting.Redis.Username,
		Password:       runtime.Setting.Redis.Password,
		MaxActiveConns: runtime.Setting.Redis.MaxActive,
	})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		panic(fmt.Errorf("failed to connect to redis: %w", err))
	}

	// 4. 启动路由
	internal.StartRouter(orm, rdb)
}
