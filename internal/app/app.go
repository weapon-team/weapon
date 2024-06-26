package app

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/middleware/requestid"

	"github.com/weapon-team/weapon/internal/app/router"
	"github.com/weapon-team/weapon/internal/sdk/engine"
)

// InitModule 初始化App模块
func InitModule(app *iris.Application, es *engine.Engines) {

	// 1.模块路由 & 中间件
	r := app.Party("/app")
	// 2.中间件
	r.Use(recover.New())
	r.Use(requestid.New(requestid.DefaultGenerator))
	// 3.依赖注入, 需配合ConfigureContainer定义路由使用
	r.RegisterDependency(es) // 注意这...不可省略
	// 4.路由分组
	router.AppUserRouter(r) // App用户
	// ...

}
