package app

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/middleware/requestid"

	"github.com/weapon-team/weapon/internal/app/router"
	"github.com/weapon-team/weapon/internal/sdk/dep"
)

// InitModule 初始化App模块
func InitModule(app *iris.Application, deps *dep.Dependency) {

	// 1.模块路由 & 中间件
	r := app.Party("/app", logger.New(logger.DefaultConfig()), recover.New(), requestid.New(requestid.DefaultGenerator))

	// 2.依赖注入, 需配合ConfigureContainer定义路由使用
	r.RegisterDependency(deps) // 注意这...不可省略

	// 3.路由分组
	router.AppUserRouter(r) // App用户
	// ...

}
