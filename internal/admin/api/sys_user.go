package api

import (
	"github.com/kataras/iris/v12"

	"github.com/weapon-team/weapon/internal/admin/model"
	adminService "github.com/weapon-team/weapon/internal/admin/service"
	appService "github.com/weapon-team/weapon/internal/app/service"
	"github.com/weapon-team/weapon/internal/sdk/engine"
	"github.com/weapon-team/weapon/internal/sdk/middleware/jwts"
	"github.com/weapon-team/weapon/internal/sdk/resp"
)

// SysUserApi 系统用户接口层
type SysUserApi struct{}

func (e SysUserApi) Party() string {
	return "user"
}

// Hello 测试接口
// path: /admin/user/hello
// param:
//
//		ctx: Iris默认可接收参数
//	 	egs: 所有第三方依赖, 如redis、orm (如不需要,不接收参数即可)
func (e SysUserApi) Hello(ctx iris.Context, egs *engine.Engines) resp.Resp {
	sysUserService := adminService.NewSysUserService(egs)
	appUserService := appService.NewAppUserService(egs)
	m := iris.Map{
		"SysUser": sysUserService.Hello(),
		"AppUser": appUserService.Hello(),
	}
	return resp.Ok(m)
}

// Login 登录
// path: /admin/user/login
func (s SysUserApi) Login(ctx iris.Context, egs *engine.Engines) resp.Resp {

	// TODO 接收参数
	us := adminService.NewSysUserService(egs)
	user := us.Login()
	token, err := jwts.GenerateToken(jwts.JwtClaims{User: user.Nickname})
	if err != nil {
		return resp.Error(1000, "", err.Error())
	}
	r := iris.Map{}
	r["user"] = user
	r["token"] = token
	return resp.Ok(r)
}

// Create
// path: /admin/user/create
func (e SysUserApi) Create(ctx iris.Context, egs *engine.Engines) resp.Resp {

	sysUserService := adminService.NewSysUserService(egs)
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
