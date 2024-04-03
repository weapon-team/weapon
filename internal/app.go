package internal

import (
	"github.com/kataras/iris/v12"
	"github.com/redis/go-redis/v9"
	"xorm.io/xorm"

	"github.com/weapon-team/weapon/internal/admin"
	"github.com/weapon-team/weapon/internal/app"
	"github.com/weapon-team/weapon/internal/sdk/middleware/jwts"
	"github.com/weapon-team/weapon/internal/sdk/resp"
	"github.com/weapon-team/weapon/internal/sdk/runtime"
)

// StartRouter 启动路由
func StartRouter(orm *xorm.Engine, rdb *redis.Client) {

	// 1. Iris
	iApp := iris.New()

	// 2. 错误处理
	iApp.OnAnyErrorCode(func(ctx iris.Context) {
		//ctx.HandlerFileLine()
		resp.Ok(ctx, resp.Resp{Code: ctx.GetStatusCode(), Data: "", Msg: ""})
	})

	// 3. 初始化Jwt
	jwts.InitJwt()

	// 4. 初始化admin模块
	admin.InitModule(iApp, orm, rdb)
	app.InitModule(iApp, orm, rdb)
	// ...

	// 4. 启动
	if err := iApp.Run(iris.Addr(":" + runtime.Setting.Port)); err != nil {
		panic(err)
	}
}
