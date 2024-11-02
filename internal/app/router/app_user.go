package router

import (
	"github.com/kataras/iris/v12"

	"github.com/weapon-team/weapon/internal/app/api"
)

type AppUserRouter struct {
	appUserApi *api.AppUserApi
}

func NewAppUserRouter(appUserApi *api.AppUserApi) *AppUserRouter {
	return &AppUserRouter{appUserApi}
}

// Register 注册路由
func (e *AppUserRouter) Register(party iris.Party) {
	party.Party("/user").ConfigureContainer(func(c *iris.APIContainer) {
		c.Get("/hello", e.appUserApi.Hello)
	})
}

func (*AppUserRouter) RegisterWithMiddleware(_ iris.Party) {

}
