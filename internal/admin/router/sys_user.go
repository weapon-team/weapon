package router

import (
	"github.com/kataras/iris/v12"

	"github.com/weapon-team/weapon/internal/admin/api"
	"github.com/weapon-team/weapon/internal/sdk/engine"
	"github.com/weapon-team/weapon/internal/sdk/middleware/casbin"
	"github.com/weapon-team/weapon/internal/sdk/middleware/jwt"
)

// SysUserRouter 用户路由组
func SysUserRouter(group iris.Party, ens *engine.Engines) {

	var suApi api.SysUserApi

	// 不需要鉴权的路由 --------------------------
	authGroup := group.Party("/auth")
	authGroup.ConfigureContainer(func(c *iris.APIContainer) {
		c.Get("/login", suApi.Login)
	})

	// 需要鉴权的路由 --------------------------
	userGroup := group.Party("/user", jwts.JwtMiddleware(), casbins.Interceptor(ens.Casbin()))
	userGroup.ConfigureContainer(func(c *iris.APIContainer) {
		c.Get("/hello", suApi.Hello)
		c.Get("/create", suApi.Create)
	})

}
