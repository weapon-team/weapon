package router

import (
	casbin2 "github.com/casbin/casbin/v2"
	"github.com/kataras/iris/v12"

	"github.com/weapon-team/weapon/internal/admin/api"
	"github.com/weapon-team/weapon/internal/sdk/middleware/casbin"
	"github.com/weapon-team/weapon/internal/sdk/middleware/jwts"
)

// SysUserRouter 用户路由组
func SysUserRouter(group iris.Party, e *casbin2.Enforcer) {

	var suApi api.SysUserApi

	//useJwtRoutes := func(c *iris.APIContainer) {
	//	c.Get("/hello", suApi.Hello)
	//	c.Get("/create", suApi.Create)
	//}
	//unUseJwtRoutes := func(c *iris.APIContainer) {
	//	c.Get("/login", suApi.Login)
	//}
	p1 := group.Party("/user")
	// 不使用jwt验证
	p1.ConfigureContainer(func(c *iris.APIContainer) {
		c.Get("/login", suApi.Login)
	})

	// 使用jwt验证
	p.ConfigureContainer(func(c *iris.APIContainer) {
		c.Get("/hello", suApi.Hello)
		c.Get("/create", suApi.Create)
	}).Use(jwts.JwtMiddleware(), casbin.CasbinInterceptor(e))

}
