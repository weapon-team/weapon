package router

import (
	"github.com/kataras/iris/v12"

	"github.com/weapon-team/weapon/internal/admin/api"
	"github.com/weapon-team/weapon/internal/sdk/engine"
	"github.com/weapon-team/weapon/internal/sdk/middleware"
)

type CaptchaRouter struct {
	capApi *api.CaptchaApi
}

func NewCaptchaRouter(capApi *api.CaptchaApi) *CaptchaRouter {
	return &CaptchaRouter{capApi}
}

// Register 注册路由 (无中间件)
func (s *CaptchaRouter) Register(party iris.Party) {

	party.Party("/captcha").ConfigureContainer(func(c *iris.APIContainer) {
		c.Get("/hello", s.capApi.Hello)
	})
}

// RegisterWithMiddleware 注册路由 (有中间件)
func (s *CaptchaRouter) RegisterWithMiddleware(party iris.Party, deps *engine.Engines) {

	party.Party("/captcha", middleware.JwtMiddleware(), middleware.PermissionInterceptor(deps.Casbin()))
	party.ConfigureContainer(func(c *iris.APIContainer) {
		c.Get("/image", s.capApi.Image)
	})
}
