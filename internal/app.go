package internal

import (
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"

	"github.com/weapon-team/weapon/internal/admin"
	"github.com/weapon-team/weapon/internal/app"
	"github.com/weapon-team/weapon/internal/sdk/engine"
	"github.com/weapon-team/weapon/internal/sdk/middleware"
	"github.com/weapon-team/weapon/internal/sdk/runtime"
	"github.com/weapon-team/weapon/pkg/logs"
)

// StartRouter 启动路由
func StartRouter(deps *engine.Engines) {

	// 1. Iris
	a := iris.New()
	a.Use(middleware.Logger())
	a.Validator = validator.New()
	a.OnAnyErrorCode(func(ctx iris.Context) {
		logs.Error().Msgf("[Iris] Request [%v][%v] %v %v, %v", ctx.Method(), ctx.GetStatusCode(), ctx.RemoteAddr(), ctx.Path(), ctx.GetErr())
	})
	// 2. 初始化Jwt
	middleware.InitJwt()

	// 3. 初始化admin模块
	admin.InitModule("/admin", a, deps)
	app.InitModule("/app", a, deps)

	// 4. 同步模型到数据库
	admin.SyncModelToTable(deps.Orm())

	// 5. 启动
	if err := a.Run(iris.Addr(":" + runtime.Setting.Port)); err != nil {
		panic(err)
	}
}
