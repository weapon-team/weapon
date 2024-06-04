package router

import (
	"github.com/kataras/iris/v12"

	"github.com/weapon-team/weapon/internal/admin/api"
	"github.com/weapon-team/weapon/internal/sdk/engine"
	"github.com/weapon-team/weapon/internal/sdk/middleware/casbins"
	"github.com/weapon-team/weapon/internal/sdk/middleware/jwts"
)

// SysRoleRouter 角色路由组
func SysRoleRouter(group iris.Party, ens *engine.Engines) {

	var srApi api.SysRoleApi

	// 不需要鉴权的路由 --------------------------
	group.Party("/role").ConfigureContainer(func(c *iris.APIContainer) {
		c.Get("/hello", srApi.Hello)
	})

	// 需要鉴权的路由 --------------------------
	group.Party("/role", jwts.JwtMiddleware(), casbins.Interceptor(ens.Casbin())).ConfigureContainer(func(c *iris.APIContainer) {

	})

}
