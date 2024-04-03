package router

import (
	"github.com/kataras/iris/v12"

	"github.com/weapon-team/weapon/internal/admin/api"
	"github.com/weapon-team/weapon/internal/sdk/middleware/jwts"
)

// SysUserRouter 用户路由组
func SysUserRouter(group iris.Party) {
	var suApi api.SysUserApi

	p := group.Party("/user")

	// 使用中间件
	p.ConfigureContainer(func(c *iris.APIContainer) {
		c.Get("/hello", suApi.Hello)
		c.Get("/create", suApi.Create)
	}).Use(jwts.JwtMiddleware())

	// 不使用中间件
	p.ConfigureContainer(func(c *iris.APIContainer) {
		c.Get("/login", suApi.Login)
	})

}
