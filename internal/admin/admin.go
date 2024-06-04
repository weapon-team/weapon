package admin

import (
	"github.com/kataras/iris/v12"
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
//	  	ens: 依赖
func InitModule(app *iris.Application, ens *engine.Engines) {

	// 1.模块路由
	r := app.Party("/admin")

	// 2.中间件
	r.Use(recover.New())
	r.Use(requestid.New(requestid.DefaultGenerator))

	// 3.依赖注入, 需配合ConfigureContainer定义路由使用
	r.RegisterDependency(ens)

	// 4.路由分组
	router.CaptchaRouter(r, ens) // 验证码路由
	router.CommonRouter(r, ens)  // 通用路由
	router.SysUserRouter(r, ens) // 系统用户路由
	router.SysRoleRouter(r, ens) // 系统角色路由
	// ...

}
