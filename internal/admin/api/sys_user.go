package api

import (
	"github.com/kataras/iris/v12"
	"xorm.io/xorm"

	"github.com/weapon-team/weapon/internal/admin/model"
	"github.com/weapon-team/weapon/internal/admin/service"
	"github.com/weapon-team/weapon/internal/sdk/resp"
)

// SysUserApi 系统用户API层
type SysUserApi struct{}

// Hello 测试接口
// path: /admin/user/hello
func (e SysUserApi) Hello(ctx iris.Context, orm *xorm.Engine) {

	sysUserService := service.NewSysUserService(orm)
	resp.Ok(ctx, sysUserService.Hello())
}

// Create
// path: /admin/user/create
func (e SysUserApi) Create(ctx iris.Context, orm *xorm.Engine) {

	sysUserService := service.NewSysUserService(orm)
	su := model.SysUser{
		Username: "Im Jordan",
		Nickname: "JJ",
		Phone:    "110",
	}
	ok := sysUserService.Create(&su)
	if !ok {
		resp.Error(ctx, 1000, "", "create error")
		return
	}

	resp.Ok(ctx, su)
}
