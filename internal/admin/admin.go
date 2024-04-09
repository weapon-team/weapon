package admin

import (
	"github.com/casbin/casbin/v2"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/middleware/requestid"

	"github.com/weapon-team/weapon/internal/admin/router"
	"github.com/weapon-team/weapon/internal/sdk/engine"
	"github.com/weapon-team/weapon/internal/sdk/middleware/jwts"
)

// InitModule 初始化admin后台模块
//
// param:
//
//		app: iris引擎
//	  	es: 依赖
func InitModule(app *iris.Application, egs *engine.Engines, e *casbin.Enforcer) {

	// 1.模块路由
	r := app.Party("/admin")
	// 2.中间件
	r.Use(logger.New(logger.DefaultConfig()))
	r.Use(recover.New())
	r.Use(requestid.New(requestid.DefaultGenerator))
	r.Use(jwts.JwtMiddleware())
	// 3.依赖注入, 需配合ConfigureContainer定义路由使用
	r.RegisterDependency(egs)
	// 4.路由分组
	router.SysUserRouter(r, e) // 系统用户路由
	router.SysRoleRouter(r)    // 系统角色路由
	// ...

}
