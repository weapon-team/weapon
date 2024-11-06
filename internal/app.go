package internal

import (
	"fmt"

	"github.com/kataras/iris/v12"
	"xorm.io/xorm"

	"github.com/weapon-team/weapon/internal/admin"
	"github.com/weapon-team/weapon/internal/admin/model"
	"github.com/weapon-team/weapon/internal/app"
	"github.com/weapon-team/weapon/internal/sdk/engine"
	"github.com/weapon-team/weapon/internal/sdk/middleware"
	"github.com/weapon-team/weapon/internal/sdk/resp"
	"github.com/weapon-team/weapon/internal/sdk/runtime"
)

// StartRouter 启动路由
func StartRouter(deps *engine.Engines) {

	// 1. Iris
	iApp := iris.New()

	// 2. 错误处理
	iApp.OnAnyErrorCode(func(ctx iris.Context) {
		resp.OkCtx(ctx, resp.Resp{Code: ctx.GetStatusCode(), Data: "", Msg: ""})
	})
	iApp.Use(middleware.Logger())

	// 3. 初始化Jwt
	middleware.InitJwt()

	// 4. 初始化admin模块
	admin.InitModule("/admin", iApp, deps)
	app.InitModule("/app", iApp, deps)

	// 5. 同步模型到数据库
	SyncModelToTable(deps.Orm())

	// 6. 启动
	if err := iApp.Run(iris.Addr(":" + runtime.Setting.Port)); err != nil {
		panic(err)
	}
}

// SyncModelToTable 同步模型到数据库
func SyncModelToTable(orm *xorm.Engine) {
	err := orm.Sync2(
		new(model.SysDept),
		new(model.SysDict),
		new(model.SysOption),
		new(model.SysRole),
		new(model.SysRoleDept),
		new(model.SysRoleMenu),
		new(model.SysUser),
		new(model.SysUserRole),
	)
	if err != nil {
		panic(fmt.Errorf("sync model to table error: %v", err))
	}
}
