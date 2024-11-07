package admin

import (
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/middleware/requestid"
	"xorm.io/xorm"

	"github.com/weapon-team/weapon/internal/admin/model"
	"github.com/weapon-team/weapon/internal/admin/router"
	"github.com/weapon-team/weapon/internal/sdk/engine"
	"github.com/weapon-team/weapon/internal/sdk/web"
)

// InitModule 初始化admin后台模块
func InitModule(path string, app *iris.Application, deps *engine.Engines) {

	// 1.模块路由
	r := app.Party(path)
	// 2.中间件
	r.Use(recover.New(), requestid.New(requestid.DefaultGenerator))

	// 3.注册所有路由
	web.RegisterRouters(r, router.AllAdminRouters(deps), deps)
	// ...

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
