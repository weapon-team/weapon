package admin

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/middleware/requestid"

	"github.com/weapon-team/weapon/internal/admin/router"
)

func InitModule(app *iris.Application, dependencies ...interface{}) {

	// 模块路由 & 中间件
	r := app.Party("/admin", logger.New(logger.DefaultConfig()), recover.New(), requestid.New(requestid.DefaultGenerator))

	// 依赖注入, 需配合ConfigureContainer定义路由使用
	r.RegisterDependency(dependencies...) // 注意这...不可省略

	router.SysUserRouter(r) // 系统用户路由
	// ...

}
