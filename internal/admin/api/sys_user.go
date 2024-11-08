package api

import (
	"github.com/kataras/iris/v12"

	"github.com/weapon-team/weapon/internal/admin/helper"
	"github.com/weapon-team/weapon/internal/admin/model"
	"github.com/weapon-team/weapon/internal/admin/service"
	"github.com/weapon-team/weapon/internal/sdk/base"
	"github.com/weapon-team/weapon/internal/sdk/middleware"
	"github.com/weapon-team/weapon/internal/sdk/resp"
)

// SysUserApi 系统用户接口层
type SysUserApi struct {
	*base.Api
	sysUserService *service.SysUserService
}

func NewSysUserApi(baseApi *base.Api, sysUserService *service.SysUserService) *SysUserApi {
	return &SysUserApi{baseApi, sysUserService}
}

// Hello 测试接口
// path: /admin/user/hello
func (e *SysUserApi) Hello(_ iris.Context) resp.Resp {
	m := iris.Map{
		"SysUser": e.sysUserService.Hello(),
	}
	return resp.OK(m)
}

// Login 登录
// path: /admin/user/login
func (e *SysUserApi) Login(ctx iris.Context) resp.Resp {

	var param helper.LoginParam
	if err := e.ReadBody(ctx, &param); err != nil {
		return resp.Error(iris.StatusBadRequest, err.Error())
	}
	user, err := e.sysUserService.Login(param)
	if err != nil {
		return resp.Error(iris.StatusBadRequest, err.Error())
	}
	claims := middleware.JwtClaims{
		Uid:      user.Id,
		Username: user.Username,
		Role:     "admin",
	}
	token, err := middleware.GenerateToken(claims)
	if err != nil {
		return resp.Error(iris.StatusBadRequest, err.Error())
	}
	r := iris.Map{
		"user":  user,
		"token": token,
	}
	return resp.OK(r)
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
		return resp.Error(iris.StatusBadRequest, "create error")
	}
	return resp.OK(su)
}
