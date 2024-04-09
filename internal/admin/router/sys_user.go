package router

import (
	"github.com/casbin/casbin/v2"
	"github.com/kataras/iris/v12"

	"github.com/weapon-team/weapon/internal/admin/api"
	"github.com/weapon-team/weapon/internal/sdk/middleware/jwts"
)

// SysUserRouter 用户路由组
func SysUserRouter(group iris.Party, e *casbin.Enforcer) {

	var suApi api.SysUserApi

	p := group.Party("/user")

	// 正常使用中间件
	p.ConfigureContainer(func(c *iris.APIContainer) {
		c.Get("/hello", suApi.Hello)
		c.Get("/create", suApi.Create)
	})

	// 去除某些中间件
	p.RemoveHandler(jwts.JwtMiddleware())
	p.ConfigureContainer(func(c *iris.APIContainer) {
		c.Get("/login", suApi.Login)
	})

}
