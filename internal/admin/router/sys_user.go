package router

import (
	"github.com/kataras/iris/v12"

	"github.com/weapon-team/weapon/internal/admin/api"
)

type SysUserRouter struct {
	userApi *api.SysUserApi
}

func NewSysUserRouter(userApi *api.SysUserApi) *SysUserRouter {
	return &SysUserRouter{userApi}
}

// Register 注册路由 (无中间件)
func (s *SysUserRouter) Register(party iris.Party) {
	authGroup := party.Party("/user")
	authGroup.ConfigureContainer(func(c *iris.APIContainer) {
		c.Get("/login", s.userApi.Login)
	})
}

// RegisterWithMiddleware 注册路由 (有中间件)
func (s *SysUserRouter) RegisterWithMiddleware(party iris.Party) {
	//userGroup := party.Party("/user", casbins.JwtMiddleware(), casbins.PermissionInterceptor(s.Casbin()))
	userGroup := party.Party("/user")
	userGroup.ConfigureContainer(func(c *iris.APIContainer) {
		c.Get("/hello", s.userApi.Hello)
		c.Get("/create", s.userApi.Create)
	})
}
