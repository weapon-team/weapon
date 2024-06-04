package router

import (
	"github.com/kataras/iris/v12"

	"github.com/weapon-team/weapon/internal/admin/api"
	"github.com/weapon-team/weapon/internal/sdk/engine"
	"github.com/weapon-team/weapon/internal/sdk/middleware/casbins"
	"github.com/weapon-team/weapon/internal/sdk/middleware/jwts"
)

// CommonRouter 通用路由组
func CommonRouter(group iris.Party, ens *engine.Engines) {

	var cApi api.CommonApi

	// 不需jwt鉴权和casbin权限验证
	group.Party("/common").ConfigureContainer(func(c *iris.APIContainer) {
		c.Get("/hello", cApi.Hello)
		c.Get("/dict/option", cApi.DictOption)
	})

	// 需要鉴权的接口
	group.Party("/common", jwts.JwtMiddleware(), casbins.Interceptor(ens.Casbin())).ConfigureContainer(func(c *iris.APIContainer) {

	})

}
