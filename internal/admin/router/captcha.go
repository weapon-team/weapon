package router

import (
	"github.com/kataras/iris/v12"

	"github.com/weapon-team/weapon/internal/admin/api"
	"github.com/weapon-team/weapon/internal/sdk/engine"
	"github.com/weapon-team/weapon/internal/sdk/web"
)

type CaptchaRouter struct {
	capApi *api.CaptchaApi
}

var _ web.IRouter = (*CaptchaRouter)(nil)

func NewCaptchaRouter(capApi *api.CaptchaApi) *CaptchaRouter {
	return &CaptchaRouter{capApi}
}

// Register 注册路由 (无中间件)
func (s *CaptchaRouter) Register(party iris.Party) {
	party.Party("/captcha").ConfigureContainer(func(c *iris.APIContainer) {
		c.Get("/hello", s.capApi.Hello)
		c.Get("/", s.capApi.Captcha)
	})
}

// RegisterWithMiddleware 注册路由 (有中间件)
func (s *CaptchaRouter) RegisterWithMiddleware(party iris.Party, deps *engine.Engines) {
	//jwtAuth, permissionAuth := middleware.JwtMiddleware(), middleware.PermissionInterceptor(deps.Casbin())
	//party.Party("/captcha", jwtAuth, permissionAuth).ConfigureContainer(func(c *iris.APIContainer) {
	//	c.Get("/image", s.capApi.Image)
	//})
}
