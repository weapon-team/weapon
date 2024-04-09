package router

import (
	"github.com/casbin/casbin/v2"
	"github.com/kataras/iris/v12"

	"github.com/weapon-team/weapon/internal/admin/api"
	"github.com/weapon-team/weapon/internal/sdk/middleware/casbins"
	"github.com/weapon-team/weapon/internal/sdk/middleware/jwts"
)

// SysUserRouter 用户路由组
func SysUserRouter(group iris.Party, e *casbin.SyncedEnforcer) {

	var suApi api.SysUserApi

	// user
	userGroup := group.Party("/user", jwts.JwtMiddleware(), casbins.Interceptor(e))
	userGroup.ConfigureContainer(func(c *iris.APIContainer) {
		c.Get("/hello", suApi.Hello)
		c.Get("/create", suApi.Create)
	})

	// auth - 不需jwt鉴权和casbin权限验证
	authGroup := group.Party("/auth")
	authGroup.ConfigureContainer(func(c *iris.APIContainer) {
		c.Get("/login", suApi.Login)
	})

}
