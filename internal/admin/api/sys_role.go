package api

import (
	"github.com/kataras/iris/v12"

	"github.com/weapon-team/weapon/internal/sdk/resp"
)

// SysRoleApi 系统角色接口层
type SysRoleApi struct {
}

func NewSysRoleApi() *SysRoleApi {
	return &SysRoleApi{}
}

// Hello
// path: /admin/role/hello
func (s *SysRoleApi) Hello(_ iris.Context) resp.Resp {
	return resp.Ok("Hello SysRole!")
}
