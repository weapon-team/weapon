package web

import (
	"github.com/kataras/iris/v12"

	"github.com/weapon-team/weapon/internal/sdk/engine"
)

// IRouter 路由接口
type IRouter interface {
	Register(r iris.Party)
	RegisterWithMiddleware(r iris.Party, deps *engine.Engines)
}

// RegisterRouters 注册路由
func RegisterRouters(party iris.Party, routers []IRouter, deps *engine.Engines) {
	for _, router := range routers {
		router.Register(party)
		router.RegisterWithMiddleware(party, deps)
	}
}
