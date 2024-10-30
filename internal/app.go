package internal

import (
	"github.com/casbin/casbin/v2"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/redis/go-redis/v9"
	"xorm.io/xorm"

	"github.com/weapon-team/weapon/internal/admin"
	"github.com/weapon-team/weapon/internal/app"
	"github.com/weapon-team/weapon/internal/sdk/engine"
	"github.com/weapon-team/weapon/internal/sdk/middleware"
	"github.com/weapon-team/weapon/internal/sdk/resp"
	"github.com/weapon-team/weapon/internal/sdk/runtime"
)

// StartRouter 启动路由
func StartRouter(orm *xorm.Engine, rdb *redis.Client, e *casbin.SyncedEnforcer) {

	// 1. Iris
	iApp := iris.New()

	// 2. 错误处理
	iApp.OnAnyErrorCode(func(ctx iris.Context) {
		//ctx.HandlerFileLine()
		resp.OkCtx(ctx, resp.Resp{Code: ctx.GetStatusCode(), Data: "", Msg: ""})
	})

	iApp.Use(logger.New(middleware.MyLogConfig()))

	// 3. 初始化Jwt
	middleware.InitJwt()

	// 4. 组装所有引擎
	deps := engine.NewEngines(orm, rdb, e)

	// 5. 初始化admin模块
	admin.InitModule(iApp, deps)
	app.InitModule(iApp, deps)
	// 6. 同步模型到数据库
	syncModels(orm)
	// 7. 启动
	if err := iApp.Run(iris.Addr(":" + runtime.Setting.Port)); err != nil {
		panic(err)
	}
}

// 同步模型到数据库
func syncModels(orm *xorm.Engine) {
	//adminModels := []any{
	//	new(model.GenConfig), new(model.GenFieldConfig), new(model.SysAnnouncement), new(model.SysDept),
	//	new(model.SysDict), new(model.SysDictItem), new(model.SysFile), new(model.SysLog),
	//	new(model.SysMenu), new(model.SysMessage), new(model.SysMessageUser), new(model.SysOption),
	//	new(model.SysRole), new(model.SysRoleDept), new(model.SysRoleMenu), new(model.SysStorage),
	//	new(model.SysUser), new(model.SysUserRole), new(model.SysUserSocial),
	//}
	//err := orm.Sync2(adminModels)
	//if err != nil {
	//	panic(err)
	//}
}
