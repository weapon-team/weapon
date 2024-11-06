package app

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/middleware/requestid"

	"github.com/weapon-team/weapon/internal/app/router"
	"github.com/weapon-team/weapon/internal/sdk/engine"
	"github.com/weapon-team/weapon/internal/sdk/web"
)

// InitModule 初始化App模块
func InitModule(path string, app *iris.Application, deps *engine.Engines) {

	// 1.模块路由
	r := app.Party(path)

	// 2.中间件
	r.Use(recover.New(), requestid.New(requestid.DefaultGenerator))

	// 3.路由分组
	web.RegisterRouters(app, router.AllAppRouters(deps), deps)
	// ...
}
