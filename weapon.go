package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/weapon-team/weapon/internal"
	engine "github.com/weapon-team/weapon/internal/sdk/engine"
	"github.com/weapon-team/weapon/internal/sdk/runtime"
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
	logs.Init(runtime.Setting.Log.Level, runtime.Setting.Log.Path, runtime.Setting.Log.Interval)

	// 3. 初始化所有依赖
	deps, err := engine.InitEngines()
	if err != nil {
		panic(err)
	}
	// 4. 启动路由
	internal.StartRouter(deps)
}
