package api

import (
	"github.com/kataras/iris/v12"

	"github.com/weapon-team/weapon/internal/admin/service"
	"github.com/weapon-team/weapon/internal/sdk/base"
	"github.com/weapon-team/weapon/internal/sdk/resp"
)

// CaptchaApi 验证码接口层
type CaptchaApi struct {
	*base.Api
	captchaService *service.CaptchaService
}

func NewCaptchaApi(baseApi *base.Api, captchaService *service.CaptchaService) *CaptchaApi {
	return &CaptchaApi{baseApi, captchaService}
}

// Hello 测试接口
// path: /captcha/hello
func (e *CaptchaApi) Hello(_ iris.Context) resp.Resp {
	return resp.OK("Hello Captcha API !")
}

// Captcha 图形验证码
// path: /captcha
func (s *CaptchaApi) Captcha() resp.Resp {
	id, b64s, err := s.captchaService.Generate()
	if err != nil {
		return resp.Error(iris.StatusBadRequest, err.Error())
	}
	return resp.OK(iris.Map{"id": id, "b64s": b64s})
}
