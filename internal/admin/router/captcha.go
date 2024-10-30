package router

import (
	"github.com/kataras/iris/v12"

	"github.com/weapon-team/weapon/internal/admin/api"
	"github.com/weapon-team/weapon/internal/sdk/engine"
	"github.com/weapon-team/weapon/internal/sdk/middleware"
)

type CaptchaRouter struct {
	*engine.Engines
}

func NewCaptchaRouter(ens *engine.Engines) *CaptchaRouter {
	return &CaptchaRouter{
		ens,
	}
}

// Register 注册路由 (无中间件)
func (s *CaptchaRouter) Register(party iris.Party) {
	captchaApi := api.NewCaptchaApi(s.Engines)
	party.Party("/captcha").ConfigureContainer(func(c *iris.APIContainer) {
		c.Get("/hello", captchaApi.Hello)
	})
}

// RegisterWithMiddleware 注册路由 (有中间件)
func (s *CaptchaRouter) RegisterWithMiddleware(party iris.Party) {
	captchaApi := api.NewCaptchaApi(s.Engines)
	party.Party("/captcha", middleware.JwtMiddleware(), middleware.PermissionInterceptor(s.Casbin())).ConfigureContainer(func(c *iris.APIContainer) {
		c.Get("/image", captchaApi.Image)
	})
}
