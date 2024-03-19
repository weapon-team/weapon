package api

import (
	"github.com/kataras/iris/v12"
	"github.com/weapon-team/weapon/internal/admin/model"
	"github.com/weapon-team/weapon/internal/admin/service"
	"github.com/weapon-team/weapon/internal/sdk/resp"
	"xorm.io/xorm"
)

type SysUserApi struct{}

// Hello 测试接口
// path: /admin/user/hello
func (e SysUserApi) Hello(ctx iris.Context, orm *xorm.Engine) {

	// ctx.Params()

	sysUserService := service.NewSysUserService(orm)

	resp.Ok(ctx, sysUserService.Hello())
}

func (e SysUserApi) Create(ctx iris.Context, orm *xorm.Engine) {

	sysUserService := service.NewSysUserService(orm)
	su := model.SysUser{
		Name:     "Im Jordan",
		NickName: "JJ",
		Phone:    "110",
	}
	ok := sysUserService.Create(&su)
	if !ok {
		resp.Error(ctx, 1000, "", "create error")
		return
	}

	resp.Ok(ctx, su)
}
