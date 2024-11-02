//go:build wireinject
// +build wireinject

package router

import (
	"github.com/google/wire"

	"github.com/weapon-team/weapon/internal/admin/api"
	"github.com/weapon-team/weapon/internal/admin/service"
	"github.com/weapon-team/weapon/internal/sdk/engine"
	"github.com/weapon-team/weapon/internal/sdk/web"
)

// 所有的Admin service都需要添加到这里
var allAdminService = wire.NewSet(
	service.NewSysUserService,
	service.NewSysRoleService,
	service.NewSysOptionService,
)

// 所有的Admin api都需要添加到这里
var allAdminApi = wire.NewSet(
	api.NewSysUserApi,
	api.NewSysRoleApi,
	api.NewCaptchaApi,
	api.NewCommonApi,
)

// 所有的Admin Router都需要添加到这里
var allAdminRouter = wire.NewSet(
	NewSysUserRouter,
	NewSysRoleRouter,
	NewCommonRouter,
	NewCaptchaRouter,
)

// 每添加一个Router都在这里添加一个参数并返回
func buildAllAdminIRouter(
	userRouter *SysUserRouter,
	roleRouter *SysRoleRouter,
	commonRouter *CommonRouter,
	captchaRouter *CaptchaRouter,
) []web.IRouter {
	return []web.IRouter{
		userRouter,
		roleRouter,
		commonRouter,
		captchaRouter,
	}
}

// AllAdminRouters 初始化所有路由并返回 []web.IRouter
func AllAdminRouters(deps *engine.Engines) []web.IRouter {
	wire.Build(allAdminService, allAdminApi, allAdminRouter, buildAllAdminIRouter) // 使用 RouterSet 生成所有路由
	return nil                                                                     // wire.Build 会替代这个返回值
}
