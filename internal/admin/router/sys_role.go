package router

import (
	"github.com/kataras/iris/v12"

	"github.com/weapon-team/weapon/internal/admin/api"
)

// SysRoleRouter 角色路由组
func SysRoleRouter(group iris.Party) {

	var srApi api.SysRoleApi

	group.Party("/role").ConfigureContainer(func(c *iris.APIContainer) {
		c.Get("/hello", srApi.Hello)
	}, func(c *iris.APIContainer) {
		//c.Get("/login", suApi.Login).RemoveHandler(jwts.JwtMiddleware(), casbins.Interceptor(e))
	})
}
