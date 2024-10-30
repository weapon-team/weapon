package api

import (
	"github.com/kataras/iris/v12"

	"github.com/weapon-team/weapon/internal/admin/service"
	"github.com/weapon-team/weapon/internal/sdk/engine"
	"github.com/weapon-team/weapon/internal/sdk/resp"
)

// AppUserApi App用户Api层
type AppUserApi struct{}

func NewAppUserApi() *AppUserApi {
	return &AppUserApi{}
}

// Hello
// path: /app/user/hello
func (api AppUserApi) Hello(_ iris.Context, egs *engine.Engines) resp.Resp {

	sus := service.NewSysUserService(egs)
	return resp.Ok(sus.Hello())
}
