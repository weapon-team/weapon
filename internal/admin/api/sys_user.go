package api

import (
	"github.com/kataras/iris/v12"

	"github.com/weapon-team/weapon/internal/admin/model"
	adminService "github.com/weapon-team/weapon/internal/admin/service"
	appService "github.com/weapon-team/weapon/internal/app/service"
	"github.com/weapon-team/weapon/internal/sdk/dep"
	"github.com/weapon-team/weapon/internal/sdk/middleware/jwts"
	"github.com/weapon-team/weapon/internal/sdk/resp"
)

// SysUserApi 系统用户API层
type SysUserApi struct{}

// Hello 测试接口
// path: /admin/user/hello
func (e SysUserApi) Hello(ctx iris.Context, deps *dep.Dependency) resp.Resp {
	sysUserService := adminService.NewSysUserService(deps)
	appUserService := appService.NewAppUserService(deps)
	m := iris.Map{
		"SysUser": sysUserService.Hello(),
		"AppUser": appUserService.Hello(),
	}
	return resp.Resp{
		Data: m,
		Code: 200,
		Msg:  "ok",
	}
}

// Login 登录
// path: /admin/user/login
func (s SysUserApi) Login(ctx iris.Context, deps *dep.Dependency) resp.Resp {

	// TODO 接收参数
	us := adminService.NewSysUserService(deps)
	user := us.Login()
	token, err := jwts.GenerateToken(jwts.JwtClaims{User: user.Nickname})
	if err != nil {
		return resp.Error(1000, "", err.Error())
	}
	r := iris.Map{
		"user":  user,
		"token": token,
	}
	return resp.Ok(r)
}

// Create
// path: /admin/user/create
func (e SysUserApi) Create(ctx iris.Context, deps *dep.Dependency) resp.Resp {

	sysUserService := adminService.NewSysUserService(deps)
	su := model.SysUser{
		Username: "Im Jordan",
		Nickname: "JJ",
		Phone:    "110",
	}
	ok := sysUserService.Create(&su)
	if !ok {

		return resp.Error(1000, "", "create error")
	}
	return resp.Ok(su)
}
