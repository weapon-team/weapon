package api

import (
	"github.com/kataras/iris/v12"

	"github.com/weapon-team/weapon/internal/admin/helper"
	"github.com/weapon-team/weapon/internal/admin/service"
	"github.com/weapon-team/weapon/internal/sdk/base"
	"github.com/weapon-team/weapon/internal/sdk/middleware"
	"github.com/weapon-team/weapon/internal/sdk/resp"
)

// SysRoleApi 系统角色接口层
type SysRoleApi struct {
	*base.Api
	roleService *service.SysRoleService
}

func NewSysRoleApi(baseApi *base.Api, roleService *service.SysRoleService) *SysRoleApi {
	return &SysRoleApi{baseApi, roleService}
}

// Hello
// path: /admin/role/hello
func (s *SysRoleApi) Hello(_ iris.Context) resp.Resp {
	return resp.OK("Hello SysRole!")
}

// Get 获取角色详情
// path: /admin/role/:id
func (e *SysRoleApi) Get(ctx iris.Context) resp.Resp {
	id := ctx.Params().GetUintDefault("id", 0)
	if id == 0 {
		return resp.Error(iris.StatusBadRequest, "id不能为空")
	}
	role, err := e.roleService.Get(id)
	if err != nil {
		return resp.Error(iris.StatusBadRequest, err.Error())
	}
	return resp.OK(role)
}

// List 分页查询
// path: /admin/role/list
func (e *SysRoleApi) List(ctx iris.Context) resp.Resp {
	var param helper.ListRoleParam
	if err := e.ReadBody(ctx, &param); err != nil {
		return resp.Error(iris.StatusBadRequest, err.Error())
	}
	pageResp, err := e.roleService.List(param)
	if err != nil {
		return resp.Error(iris.StatusBadRequest, err.Error())
	}
	return resp.OK(pageResp)
}

// Create
// path: /admin/role/create
func (e *SysRoleApi) Create(ctx iris.Context) resp.Resp {

	var param helper.CreateRoleParam
	if err := e.ReadBody(ctx, &param); err != nil {
		return resp.Error(iris.StatusBadRequest, err.Error())
	}
	optUserId := middleware.GetUserId(ctx)
	u, err := e.roleService.Create(param, optUserId)
	if err != nil {
		return resp.Error(iris.StatusBadRequest, err.Error())
	}
	return resp.OK(u)
}

// Update 修改用户
// path: /admin/role/update
func (e *SysRoleApi) Update(ctx iris.Context) resp.Resp {

	var param helper.UpdateRoleParam
	if err := e.ReadBody(ctx, &param); err != nil {
		return resp.Error(iris.StatusBadRequest, err.Error())
	}
	optUserId := middleware.GetUserId(ctx)
	newUser, err := e.roleService.Update(param, optUserId)
	if err != nil {
		return resp.Error(iris.StatusBadRequest, err.Error())
	}
	return resp.OK(newUser)
}

// Delete 删除角色
// path: /admin/role/delete
func (e *SysRoleApi) Delete(ctx iris.Context) resp.Resp {
	uid := ctx.Params().GetUintDefault("id", 0)
	if uid == 0 {
		return resp.Error(iris.StatusBadRequest, "id不能为空")
	}
	if err := e.roleService.Delete(uid); err != nil {
		return resp.Error(iris.StatusBadRequest, err.Error())
	}
	return resp.OK()
}
