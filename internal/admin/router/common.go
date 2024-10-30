package router

import (
	"github.com/kataras/iris/v12"

	"github.com/weapon-team/weapon/internal/admin/api"
	"github.com/weapon-team/weapon/internal/sdk/engine"
	"github.com/weapon-team/weapon/internal/sdk/middleware"
)

type CommonRouter struct {
	*engine.Engines
}

func NewCommonRouter(ens *engine.Engines) *CommonRouter {
	return &CommonRouter{
		ens,
	}
}

// Register 注册路由 (无中间件)
func (s *CommonRouter) Register(party iris.Party) {
	commonApi := api.NewCommonApi(s.Engines)
	// 不需jwt鉴权和casbin权限验证
	party.Party("/common").
		ConfigureContainer(func(c *iris.APIContainer) {
			c.Get("/hello", commonApi.Hello)
			c.Get("/dict/option", commonApi.DictOption)
		})
}

// RegisterWithMiddleware 注册路由 (有中间件)
func (s *CommonRouter) RegisterWithMiddleware(party iris.Party) {
	party.Party("/common", middleware.JwtMiddleware(), middleware.PermissionInterceptor(s.Casbin())).
		ConfigureContainer(func(c *iris.APIContainer) {

		})
}
