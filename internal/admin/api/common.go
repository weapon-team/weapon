package api

import (
	"github.com/kataras/iris/v12"

	"github.com/weapon-team/weapon/internal/admin/service"
	"github.com/weapon-team/weapon/internal/sdk/resp"
)

// CommonApi 系统用户接口层
type CommonApi struct {
	optService *service.SysOptionService
}

func NewCommonApi(optService *service.SysOptionService) *CommonApi {
	return &CommonApi{optService}
}

func (e *CommonApi) Hello(_ iris.Context) resp.Resp {
	return resp.Ok("Hello Common API !")
}

// DictOption 查询参数字典
// path: /admin/common/dict/option
// param:
//
//		ctx: Iris默认可接收参数
//	 	egs: 所有第三方依赖, 如redis、orm (如不需要,不接收参数即可)
func (e *CommonApi) DictOption(_ iris.Context) resp.Resp {
	data, err := e.optService.AllOptions()
	if err != nil {
		return resp.Error(iris.StatusBadRequest, "", err.Error())
	}
	return resp.Ok(data)
}
