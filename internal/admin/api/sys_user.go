package api

import (
	"github.com/kataras/iris/v12"

	"github.com/weapon-team/weapon/internal/admin/model"
	adminService "github.com/weapon-team/weapon/internal/admin/service"
	appService "github.com/weapon-team/weapon/internal/app/service"
	"github.com/weapon-team/weapon/internal/sdk/engine"
	"github.com/weapon-team/weapon/internal/sdk/middleware"
	"github.com/weapon-team/weapon/internal/sdk/resp"
)

// SysUserApi 系统用户接口层
type SysUserApi struct {
	*engine.Engines
}

func NewSysUserApi(deps *engine.Engines) *SysUserApi {
	return &SysUserApi{deps}
}

// Hello 测试接口
// path: /admin/user/hello
func (e *SysUserApi) Hello(_ iris.Context) resp.Resp {
	sysUserService := adminService.NewSysUserService(e.Engines)
	appUserService := appService.NewAppUserService(e.Engines)
	m := iris.Map{
		"SysUser": sysUserService.Hello(),
		"AppUser": appUserService.Hello(),
	}
	return resp.Ok(m)
}

// Login 登录
// path: /admin/user/login
func (e *SysUserApi) Login(_ iris.Context) resp.Resp {
	// TODO 接收参数
	us := adminService.NewSysUserService(e.Engines)
	user := us.Login()
	token, err := middleware.GenerateToken(middleware.JwtClaims{Uid: user.Id, Username: user.Username, Role: "admin"})
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
func (e *SysUserApi) Create(_ iris.Context) resp.Resp {

	sysUserService := adminService.NewSysUserService(e.Engines)
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
