package api

import (
	"github.com/kataras/iris/v12"
	"xorm.io/xorm"

	"github.com/weapon-team/weapon/internal/admin/service"
	"github.com/weapon-team/weapon/internal/sdk/resp"
)

// AppUserApi App用户Api层
type AppUserApi struct{}

// Hello
// path: /app/user/hello
func (api AppUserApi) Hello(ctx iris.Context, orm *xorm.Engine) resp.Resp {

	sus := service.NewSysUserService(orm)
	return resp.Ok(sus.Hello())
}
