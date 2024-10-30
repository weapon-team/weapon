package api

import (
	"github.com/kataras/iris/v12"

	adminService "github.com/weapon-team/weapon/internal/admin/service"
	"github.com/weapon-team/weapon/internal/sdk/engine"
	"github.com/weapon-team/weapon/internal/sdk/resp"
)

// CommonApi 系统用户接口层
type CommonApi struct {
	*engine.Engines
}

func NewCommonApi(deps *engine.Engines) *CommonApi {
	return &CommonApi{deps}
}

func (e *CommonApi) Hello(_ iris.Context, _ *engine.Engines) resp.Resp {
	return resp.Ok("Hello Common API !")
}

// DictOption 查询参数字典
// path: /admin/common/dict/option
// param:
//
//		ctx: Iris默认可接收参数
//	 	egs: 所有第三方依赖, 如redis、orm (如不需要,不接收参数即可)
func (e *CommonApi) DictOption(_ iris.Context, ens *engine.Engines) resp.Resp {

	optSve := adminService.NewSysOptionService(ens)

	data, err := optSve.AllOptions()
	if err != nil {
		return resp.Error(iris.StatusBadRequest, "", err.Error())
	}
	return resp.Ok(data)
}
