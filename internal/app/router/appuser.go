package router

import (
	"github.com/kataras/iris/v12"
	"github.com/weapon-team/weapon/internal/app/api"
)

// AppUserRouter 用户路由组
func AppUserRouter(group iris.Party) {

	var auApi api.AppUserApi
	group.Party("/user").ConfigureContainer(func(c *iris.APIContainer) {
		c.Get("/hello", auApi.Hello)
	})

}
