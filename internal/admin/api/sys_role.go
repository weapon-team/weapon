package api

import (
	"github.com/kataras/iris/v12"

	"github.com/weapon-team/weapon/internal/sdk/base"
	"github.com/weapon-team/weapon/internal/sdk/resp"
)

// SysRoleApi 系统角色接口层
type SysRoleApi struct {
	*base.Api
}

func NewSysRoleApi(baseApi *base.Api) *SysRoleApi {
	return &SysRoleApi{baseApi}
}

// Hello
// path: /admin/role/hello
func (s *SysRoleApi) Hello(_ iris.Context) resp.Resp {
	return resp.Ok("Hello SysRole!")
}
