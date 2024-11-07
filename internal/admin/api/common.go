package api

import (
	"github.com/kataras/iris/v12"

	"github.com/weapon-team/weapon/internal/admin/service"
	"github.com/weapon-team/weapon/internal/sdk/base"
	"github.com/weapon-team/weapon/internal/sdk/resp"
)

// CommonApi 系统用户接口层
type CommonApi struct {
	*base.Api
	optService *service.SysOptionService
}

func NewCommonApi(baseApi *base.Api, optService *service.SysOptionService) *CommonApi {
	return &CommonApi{Api: baseApi, optService: optService}
}

// Hello 测试接口
// path: /common/hello
func (e *CommonApi) Hello(_ iris.Context) resp.Resp {
	return resp.OK("Hello Common API !")
}

// DictOption 查询参数字典
// path: /admin/common/dict/option
func (e *CommonApi) DictOption(_ iris.Context) resp.Resp {
	data, err := e.optService.AllOptions()
	if err != nil {
		return resp.Error(iris.StatusBadRequest, err.Error())
	}
	return resp.OK(data)
}
