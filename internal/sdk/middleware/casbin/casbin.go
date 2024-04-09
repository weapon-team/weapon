package casbin

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	"github.com/kataras/iris/v12"
)

func CasbinInterceptor(e *casbin.Enforcer) iris.Handler {
	return func(c iris.Context) {

		//获取请求的URI
		obj := c.Request().RequestURI
		//获取请求方法
		act := c.Request().Method
		//获取用户的角色
		sub := "admin"

		//判断策略中是否存在
		if ok, _ := e.Enforce(sub, obj, act); ok {
			fmt.Println("恭喜您,权限验证通过")
			c.Next()
		} else {
			fmt.Println("很遗憾,权限验证没有通过")
			_ = c.StopWithJSON(200, iris.Map{"code": 2000, "data": "", "msg": "您不具备此权限"})
		}
	}
}
