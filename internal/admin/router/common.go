package router

import (
	"github.com/kataras/iris/v12"

	"github.com/weapon-team/weapon/internal/admin/api"
)

type CommonRouter struct {
	cmApi *api.CommonApi
}

func NewCommonRouter(cmApi *api.CommonApi) *CommonRouter {
	return &CommonRouter{
		cmApi,
	}
}

// Register 注册路由 (无中间件)
func (s *CommonRouter) Register(party iris.Party) {

	// 不需jwt鉴权和casbin权限验证
	party.Party("/common").
		ConfigureContainer(func(c *iris.APIContainer) {
			c.Get("/hello", s.cmApi.Hello)
			c.Get("/dict/option", s.cmApi.DictOption)
		})
}

// RegisterWithMiddleware 注册路由 (有中间件)
func (s *CommonRouter) RegisterWithMiddleware(party iris.Party) {
	//party.Party("/common", middleware.JwtMiddleware(), middleware.PermissionInterceptor(s.Casbin())).
	party.Party("/common").
		ConfigureContainer(func(c *iris.APIContainer) {

		})
}
