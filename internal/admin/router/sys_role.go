package router

import (
	"github.com/kataras/iris/v12"

	"github.com/weapon-team/weapon/internal/admin/api"
)

// SysRoleRouter 角色路由组
func SysRoleRouter(group iris.Party) {

	var srApi api.SysRoleApi
	p := group.Party("/role")

	// 使用jwt验证
	p.ConfigureContainer(func(c *iris.APIContainer) {
		c.Get("/hello", srApi.Hello)
	})

	// 不使用jwt验证
	p.RemoveHandler()
	p.ConfigureContainer(func(c *iris.APIContainer) {

	})
}
