package internal

import (
	"github.com/kataras/iris/v12"
	"github.com/weapon-team/weapon/internal/admin"
	"github.com/weapon-team/weapon/internal/app"
	"github.com/weapon-team/weapon/internal/sdk/runtime"
	"github.com/weapon-team/weapon/pkg/ternary"
	"xorm.io/xorm"
)

// StartRouter 启动路由
func StartRouter(orm *xorm.Engine) {
	// 1. Iris
	iApp := iris.New()

	// 2. 错误处理
	iApp.OnAnyErrorCode(func(ctx iris.Context) {

		var err1, err2 = ctx.GetErr(), ctx.Err()
		err := ternary.If(err1 != nil, err1, err2)
		msg := ternary.If(err != nil, err.Error(), "")

		ctx.JSON(iris.Map{"code": ctx.GetStatusCode(), "data": "", "msg": msg})
	})

	// 3. 初始化admin模块
	admin.InitModule(iApp, orm)
	app.InitModule(iApp, orm)
	// ...

	// 4. 启动
	iApp.Run(iris.Addr(":" + runtime.Setting.Port))
}
