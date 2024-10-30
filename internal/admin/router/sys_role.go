package router

import (
	"github.com/kataras/iris/v12"

	"github.com/weapon-team/weapon/internal/admin/api"
	"github.com/weapon-team/weapon/internal/sdk/engine"
	"github.com/weapon-team/weapon/internal/sdk/middleware"
)

type SysRoleRouter struct {
	*engine.Engines
}

func NewSysRoleRouter(ens *engine.Engines) *SysRoleRouter {
	return &SysRoleRouter{
		ens,
	}
}

// Register 注册路由 (无中间件)
func (s *SysRoleRouter) Register(party iris.Party) {
	sysRoleApi := api.NewSysRoleApi(s.Engines)
	// 不需jwt鉴权和casbin权限验证
	party.Party("/role").ConfigureContainer(func(c *iris.APIContainer) {
		c.Get("/hello", sysRoleApi.Hello)
	})
}

// RegisterWithMiddleware 注册路由 (有中间件)
func (s *SysRoleRouter) RegisterWithMiddleware(party iris.Party) {
	//sysRoleApi := api.NewSysRoleApi(s.Engines)
	party.Party("/role", middleware.JwtMiddleware(), middleware.PermissionInterceptor(s.Casbin())).
		ConfigureContainer(func(c *iris.APIContainer) {

		})
}
