package router

import (
	"github.com/kataras/iris/v12"

	"github.com/weapon-team/weapon/internal/app/api"
	"github.com/weapon-team/weapon/internal/sdk/engine"
)

type AppUserRouter struct {
	*engine.Engines
}

func NewAppUserRouter(ens *engine.Engines) *AppUserRouter {
	return &AppUserRouter{
		ens,
	}
}

// Register 注册路由
func (*AppUserRouter) Register(party iris.Party) {
	appUserApi := api.NewAppUserApi()
	party.Party("/user").ConfigureContainer(func(c *iris.APIContainer) {
		c.Get("/hello", appUserApi.Hello)
	})
}

func (*AppUserRouter) RegisterWithMiddleware(_ iris.Party) {

}
