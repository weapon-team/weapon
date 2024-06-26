package api

import (
	"github.com/kataras/iris/v12"

	"github.com/weapon-team/weapon/internal/sdk/engine"
	"github.com/weapon-team/weapon/internal/sdk/resp"
)

// CaptchaApi 验证码接口层
type CaptchaApi struct{}

func (e CaptchaApi) Hello(_ iris.Context, _ *engine.Engines) resp.Resp {
	return resp.Ok("Hello Captcha API !")
}

// Image 图形验证码
// path: /admin/captcha/image
func (s CaptchaApi) Image() resp.Resp {

	//TODO 生成图形验证码
	return resp.Ok("待完成")
}
