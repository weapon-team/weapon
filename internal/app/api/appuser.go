package api

import (
	"github.com/kataras/iris/v12"
	"github.com/weapon-team/weapon/internal/sdk/resp"
	"xorm.io/xorm"
)

type AppUserApi struct{}

// Hello
// path: /app/user/hello
func (api AppUserApi) Hello(ctx iris.Context, orm *xorm.Engine) {

	resp.Ok(ctx, "Hello App User!")
}
