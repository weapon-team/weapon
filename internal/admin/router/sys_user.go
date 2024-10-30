package router

import (
	"github.com/kataras/iris/v12"

	"github.com/weapon-team/weapon/internal/admin/api"
	"github.com/weapon-team/weapon/internal/sdk/engine"
	casbins "github.com/weapon-team/weapon/internal/sdk/middleware"
)

type SysUserRouter struct {
	*engine.Engines
}

func NewSysUserRouter(ens *engine.Engines) *SysUserRouter {
	return &SysUserRouter{
		ens,
	}
}

// Register 注册路由 (无中间件)
func (s *SysUserRouter) Register(party iris.Party) {
	sysUserApi := api.NewSysUserApi()
	authGroup := party.Party("/user")
	authGroup.ConfigureContainer(func(c *iris.APIContainer) {
		c.Get("/login", sysUserApi.Login)
	})
}

// RegisterWithMiddleware 注册路由 (有中间件)
func (s *SysUserRouter) RegisterWithMiddleware(party iris.Party) {
	sysUserApi := api.NewSysUserApi()
	userGroup := party.Party("/user", casbins.JwtMiddleware(), casbins.PermissionInterceptor(s.Casbin()))
	userGroup.ConfigureContainer(func(c *iris.APIContainer) {
		c.Get("/hello", sysUserApi.Hello)
		c.Get("/create", sysUserApi.Create)
	})
}
