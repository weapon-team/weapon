package admin

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/middleware/requestid"

	"github.com/weapon-team/weapon/internal/admin/router"
	"github.com/weapon-team/weapon/internal/sdk/engine"
)

// InitModule 初始化admin后台模块
//
// param:
//
//		app: iris引擎
//	  	dependencies: 依赖
func InitModule(app *iris.Application, es *engine.Engines) {

	// 1.模块路由 & 中间件
	r := app.Party("/admin", logger.New(logger.DefaultConfig()), recover.New(), requestid.New(requestid.DefaultGenerator))

	// 2.依赖注入, 需配合ConfigureContainer定义路由使用(注意...不可省略)
	r.RegisterDependency(es)

	// 3.路由分组
	router.SysUserRouter(r) // 系统用户路由
	// ...

}
