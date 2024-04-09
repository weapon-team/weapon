package casbins

import (
	"github.com/casbin/casbin/v2"
	"github.com/kataras/iris/v12"

	"github.com/weapon-team/weapon/internal/sdk/middleware/jwts"
)

// Interceptor 角色权限验证
func Interceptor(e *casbin.SyncedEnforcer) iris.Handler {
	return func(ctx iris.Context) {
		jc := jwts.GetJwtClaims(ctx)
		sub := jc.Role                  // 角色
		obj := ctx.Request().RequestURI //获取请求的URI
		act := ctx.Request().Method     //获取请求方法

		// 超级管理员拥有所有权限
		if sub == "admin" {
			ctx.Next()
			return
		}
		//判断策略中是否存在
		if ok, _ := e.Enforce(sub, obj, act); ok {
			ctx.Next()
		} else {
			_ = ctx.StopWithJSON(iris.StatusOK, iris.Map{"code": iris.StatusForbidden, "data": "", "msg": "很遗憾,权限验证没有通过"})
		}
	}
}
