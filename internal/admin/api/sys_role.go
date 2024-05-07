package api

import (
	"github.com/kataras/iris/v12"

	"github.com/weapon-team/weapon/internal/sdk/engine"
	"github.com/weapon-team/weapon/internal/sdk/resp"
)

// SysRoleApi 系统角色接口层
type SysRoleApi struct{}

// Hello
// path: /admin/role/hello
func (s SysRoleApi) Hello(ctx iris.Context, ens *engine.Engines) resp.Resp {

	return resp.Ok("Hello SysRole!")
}
