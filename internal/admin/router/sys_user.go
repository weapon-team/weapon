package router

import (
	"github.com/kataras/iris/v12"
	"github.com/weapon-team/weapon/internal/admin/api"
)

// SysUserRouter 用户路由组
func SysUserRouter(group iris.Party) {

	suApi := api.SysUserApi{}
	group.Party("/user").ConfigureContainer(func(c *iris.APIContainer) {
		c.Get("/hello", suApi.Hello)
		c.Get("/create", suApi.Create)
	})

}
