//go:build wireinject
// +build wireinject

package router

import (
	"github.com/google/wire"

	"github.com/weapon-team/weapon/internal/app/api"
	"github.com/weapon-team/weapon/internal/app/service"
	"github.com/weapon-team/weapon/internal/sdk/engine"
	"github.com/weapon-team/weapon/internal/sdk/web"
)

// 所有的service都需要添加到这里
var allAppService = wire.NewSet(
	service.NewAppUserService,
)

// 所有的api都需要添加到这里
var allAppApi = wire.NewSet(
	api.NewAppUserApi,
)

// 所有的Router都需要添加到这里
var allAppRouter = wire.NewSet(
	NewAppUserRouter,
)

// 每添加一个Router都在这里添加一个参数并返回
func buildAllAppIRouter(
	appUserRouter *AppUserRouter,
) []web.IRouter {
	return []web.IRouter{
		appUserRouter,
	}
}

// AllAppRouters 初始化所有路由并返回 []web.IRouter
func AllAppRouters(deps *engine.Engines) []web.IRouter {
	wire.Build(allAppService, allAppApi, allAppRouter, buildAllAppIRouter) // 使用 RouterSet 生成所有路由
	return nil                                                             // wire.Build 会替代这个返回值
}
