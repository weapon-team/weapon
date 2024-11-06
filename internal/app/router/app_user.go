package router

import (
	"github.com/kataras/iris/v12"

	"github.com/weapon-team/weapon/internal/app/api"
	"github.com/weapon-team/weapon/internal/sdk/engine"
)

type AppUserRouter struct {
	appUserApi *api.AppUserApi
}

func NewAppUserRouter(appUserApi *api.AppUserApi) *AppUserRouter {
	return &AppUserRouter{appUserApi}
}

// Register 注册路由
func (r *AppUserRouter) Register(party iris.Party) {
	party.Party("/user").ConfigureContainer(func(c *iris.APIContainer) {
		c.Get("/hello", r.appUserApi.Hello)
	})
}

func (r *AppUserRouter) RegisterWithMiddleware(party iris.Party, _ *engine.Engines) {
	party.Party("/user").ConfigureContainer(func(c *iris.APIContainer) {

	})
}
