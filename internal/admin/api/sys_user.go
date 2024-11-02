package api

import (
	"github.com/kataras/iris/v12"

	"github.com/weapon-team/weapon/internal/admin/model"
	"github.com/weapon-team/weapon/internal/admin/service"
	"github.com/weapon-team/weapon/internal/sdk/middleware"
	"github.com/weapon-team/weapon/internal/sdk/resp"
)

// SysUserApi 系统用户接口层
type SysUserApi struct {
	sysUserService *service.SysUserService
}

func NewSysUserApi(sysUserService *service.SysUserService) *SysUserApi {
	return &SysUserApi{sysUserService}
}

// Hello 测试接口
// path: /admin/user/hello
func (e *SysUserApi) Hello(_ iris.Context) resp.Resp {
	m := iris.Map{
		"SysUser": e.sysUserService.Hello(),
	}
	return resp.Ok(m)
}

// Login 登录
// path: /admin/user/login
func (e *SysUserApi) Login(_ iris.Context) resp.Resp {
	user := e.sysUserService.Login()
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
	su := model.SysUser{
		Username: "Im Jordan",
		Nickname: "JJ",
		Phone:    "110",
	}
	ok := e.sysUserService.Create(&su)
	if !ok {
		return resp.Error(1000, "", "create error")
	}
	return resp.Ok(su)
}
