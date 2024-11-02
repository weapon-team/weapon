package web

import (
	"github.com/kataras/iris/v12"
)

// IRouter 路由接口
type IRouter interface {
	Register(r iris.Party)
	RegisterWithMiddleware(r iris.Party)
}

// RegisterRouters 注册路由
func RegisterRouters(party iris.Party, routers []IRouter) {
	for _, router := range routers {
		router.Register(party)
		router.RegisterWithMiddleware(party)
	}
}
