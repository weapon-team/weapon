package router

import (
	"github.com/kataras/iris/v12"

	"github.com/weapon-team/weapon/internal/admin/api"
	"github.com/weapon-team/weapon/internal/sdk/engine"
	"github.com/weapon-team/weapon/internal/sdk/middleware"
)

type SysRoleRouter struct {
	roleApi *api.SysRoleApi
}

func NewSysRoleRouter(roleApi *api.SysRoleApi) *SysRoleRouter {
	return &SysRoleRouter{roleApi}
}

// Register 注册路由 (无中间件)
func (s *SysRoleRouter) Register(party iris.Party) {
	// 不需jwt鉴权和casbin权限验证
	party.Party("/role").ConfigureContainer(func(c *iris.APIContainer) {
		c.Get("/hello", s.roleApi.Hello)
	})
}

// RegisterWithMiddleware 注册路由 (有中间件)
func (s *SysRoleRouter) RegisterWithMiddleware(party iris.Party, deps *engine.Engines) {
	party.Party("/role", middleware.JwtMiddleware(), middleware.PermissionInterceptor(deps.Casbin()))
	party.ConfigureContainer(func(c *iris.APIContainer) {

	})
}
