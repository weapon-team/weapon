package api

import (
	"github.com/kataras/iris/v12"

	"github.com/weapon-team/weapon/internal/admin/helper"
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

// Get 获取用户详情
// path: /admin/user/:id
func (e *SysUserApi) Get(ctx iris.Context) resp.Resp {
	id := ctx.Params().GetInt64Default("id", 0)
	if id == 0 {
		return resp.Error(iris.StatusBadRequest, "id不能为空")
	}
	user, err := e.sysUserService.Get(id)
	if err != nil {
		return resp.Error(iris.StatusBadRequest, err.Error())
	}
	return resp.OK(user)
}

// List 分页查询
// path: /admin/user/list
func (e *SysUserApi) List(ctx iris.Context) resp.Resp {
	var param helper.ListUserParam
	if err := e.ReadBody(ctx, &param); err != nil {
		return resp.Error(iris.StatusBadRequest, err.Error())
	}
	pageResp, err := e.sysUserService.List(param)
	if err != nil {
		return resp.Error(iris.StatusBadRequest, err.Error())
	}
	return resp.OK(pageResp)
}

// Create
// path: /admin/user/create
func (e *SysUserApi) Create(ctx iris.Context) resp.Resp {

	var param helper.CreateUserParam
	if err := e.ReadBody(ctx, &param); err != nil {
		return resp.Error(iris.StatusBadRequest, err.Error())
	}
	optUserId := middleware.GetUserId(ctx)
	u, err := e.sysUserService.Create(param, optUserId)
	if err != nil {
		return resp.Error(iris.StatusBadRequest, err.Error())
	}
	return resp.OK(u)
}

// Update 修改用户
// path: /admin/user/update
func (e *SysUserApi) Update(ctx iris.Context) resp.Resp {

	var param helper.UpdateUserParam
	if err := e.ReadBody(ctx, &param); err != nil {
		return resp.Error(iris.StatusBadRequest, err.Error())
	}
	optUserId := middleware.GetUserId(ctx)
	newUser, err := e.sysUserService.Update(param, optUserId)
	if err != nil {
		return resp.Error(iris.StatusBadRequest, err.Error())
	}
	return resp.OK(newUser)
}

func (e *SysUserApi) Delete(ctx iris.Context) resp.Resp {
	uid := ctx.Params().GetInt64Default("id", 0)
	if uid == 0 {
		return resp.Error(iris.StatusBadRequest, "id不能为空")
	}
	if err := e.sysUserService.Delete(uid); err != nil {
		return resp.Error(iris.StatusBadRequest, err.Error())
	}
	return resp.OK()
}
