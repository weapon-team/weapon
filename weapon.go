package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/weapon-team/weapon/internal"
	"github.com/weapon-team/weapon/internal/sdk/runtime"
	"xorm.io/xorm"
)

func main() {

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
	if err := orm.Ping(); err != nil {
		panic(err)
	}
	orm.ShowSQL(true)

	// 3. redis缓存

	// 4. 启动路由
	internal.StartRouter(orm)
}
