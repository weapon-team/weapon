package router

import (
	"github.com/kataras/iris/v12"

	"github.com/weapon-team/weapon/internal/admin/api"
	"github.com/weapon-team/weapon/internal/sdk/middleware/jwts"
)

// SysRoleRouter 角色路由组
func SysRoleRouter(group iris.Party) {

	p := group.Party("/role")

	var srApi api.SysRoleApi

	// 使用jwt中间件
	useJwtRoutes := func(c *iris.APIContainer) {
		c.Get("/hello", srApi.Hello)
	}
	// 不使用jwt中间件
	unUseJwtRoutes := func(c *iris.APIContainer) {
	}

	// 使用jwt验证
	p.ConfigureContainer(useJwtRoutes).Use(jwts.JwtMiddleware())
	// 不使用jwt验证
	p.ConfigureContainer(unUseJwtRoutes)
}
