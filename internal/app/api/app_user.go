package api

import (
	"github.com/kataras/iris/v12"

	"github.com/weapon-team/weapon/internal/admin/service"
	"github.com/weapon-team/weapon/internal/sdk/engine"
	"github.com/weapon-team/weapon/internal/sdk/resp"
)

// AppUserApi App用户Api层
type AppUserApi struct {
	*engine.Engines
}

func NewAppUserApi(deps *engine.Engines) *AppUserApi {
	return &AppUserApi{deps}
}

// Hello
// path: /app/user/hello
func (e *AppUserApi) Hello(_ iris.Context) resp.Resp {
	sus := service.NewSysUserService(e.Engines)
	return resp.Ok(sus.Hello())
}
