package admin

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/middleware/requestid"

	"github.com/weapon-team/weapon/internal/admin/router"
	"github.com/weapon-team/weapon/internal/sdk/engine"
	"github.com/weapon-team/weapon/internal/sdk/web"
)

// InitModule 初始化admin后台模块
//
// param:
//
//		app: iris引擎
//	  	ens: 依赖
func InitModule(app *iris.Application, deps *engine.Engines) {
	var (
		recoverHandler   = recover.New()
		requestIdHandler = requestid.New(requestid.DefaultGenerator)
		//casbinHandler    = middleware.Interceptor(ens.Casbin())
		//jwtHandler       = middleware.JwtMiddleware()
	)
	// 1.模块路由
	r := app.Party("/admin")

	// 2.中间件
	r.Use(recoverHandler, requestIdHandler)

	// 3.依赖注入, 需配合ConfigureContainer定义路由使用
	r.RegisterDependency(deps)

	// 4.路由分组
	web.RegisterRouters(r,
		router.NewCaptchaRouter(deps), // 验证码路由
		router.NewSysUserRouter(deps), // 系统用户路由
		router.NewCommonRouter(deps),  // 公共路由
		router.NewSysRoleRouter(deps), // 系统角色路由
	)
	// ...

}
