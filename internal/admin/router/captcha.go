package router

import (
	"github.com/kataras/iris/v12"

	"github.com/weapon-team/weapon/internal/admin/api"
	"github.com/weapon-team/weapon/internal/sdk/engine"
	"github.com/weapon-team/weapon/internal/sdk/middleware/casbin"
	"github.com/weapon-team/weapon/internal/sdk/middleware/jwt"
)

// CaptchaRouter 验证码路由组
func CaptchaRouter(group iris.Party, ens *engine.Engines) {

	var cApi api.CaptchaApi

	// 不需jwt鉴权和casbin权限验证
	group.Party("/captcha").ConfigureContainer(func(c *iris.APIContainer) {
		c.Get("/hello", cApi.Hello)
	})

	// 需要鉴权的接口
	group.Party("/captcha", jwts.JwtMiddleware(), casbins.Interceptor(ens.Casbin())).ConfigureContainer(func(c *iris.APIContainer) {
		c.Get("/image", cApi.Image)
	})

}
