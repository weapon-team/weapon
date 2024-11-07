package router

import (
	"github.com/kataras/iris/v12"

	"github.com/weapon-team/weapon/internal/admin/api"
	"github.com/weapon-team/weapon/internal/sdk/engine"
	"github.com/weapon-team/weapon/internal/sdk/middleware"
	"github.com/weapon-team/weapon/internal/sdk/web"
)

type SysUserRouter struct {
	userApi *api.SysUserApi
}

var _ web.IRouter = (*SysUserRouter)(nil)

func NewSysUserRouter(userApi *api.SysUserApi) *SysUserRouter {
	return &SysUserRouter{userApi}
}

// Register 注册路由 (无中间件)
func (s *SysUserRouter) Register(party iris.Party) {
	party.Party("/user").ConfigureContainer(func(c *iris.APIContainer) {
		c.Get("/hello", s.userApi.Hello)
		c.Post("/login", s.userApi.Login)
	})
}

// RegisterWithMiddleware 注册路由 (有中间件)
func (s *SysUserRouter) RegisterWithMiddleware(party iris.Party, deps *engine.Engines) {
	jwtAuth, permissionAuth := middleware.JwtMiddleware(), middleware.PermissionInterceptor(deps.Casbin())
	party.Party("/user", jwtAuth, permissionAuth).ConfigureContainer(func(c *iris.APIContainer) {
		c.Get("/create", s.userApi.Create)
	})
}
