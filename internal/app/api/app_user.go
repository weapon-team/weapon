package api

import (
	"github.com/kataras/iris/v12"

	"github.com/weapon-team/weapon/internal/app/service"
	"github.com/weapon-team/weapon/internal/sdk/resp"
)

// AppUserApi App用户Api层
type AppUserApi struct {
	appUserService *service.AppUserService
}

func NewAppUserApi(appUserService *service.AppUserService) *AppUserApi {
	return &AppUserApi{appUserService}
}

// Hello
// path: /app/user/hello
func (e *AppUserApi) Hello(_ iris.Context) resp.Resp {
	return resp.OK("hello ! app user")
}
