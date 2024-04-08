package router

import (
	"github.com/kataras/iris/v12"

	"github.com/weapon-team/weapon/internal/admin/api"
	"github.com/weapon-team/weapon/internal/sdk/middleware/jwts"
)

// SysUserRouter 用户路由组
func SysUserRouter(group iris.Party) {

	p := group.Party("/user")
	var suApi api.SysUserApi

	useJwtRoutes := func(c *iris.APIContainer) {
		c.Get("/hello", suApi.Hello)
		c.Get("/create", suApi.Create)
	}
	unUseJwtRoutes := func(c *iris.APIContainer) {
		c.Get("/login", suApi.Login)
	}

	// 使用中间件
	p.ConfigureContainer(useJwtRoutes).Use(jwts.JwtMiddleware())
	// 不使用中间件
	p.ConfigureContainer(unUseJwtRoutes)
}
