package resp

import (
	"github.com/kataras/iris/v12"
)

// Resp 响应体结构
type Resp struct {
	Code int    `json:"code"` // 状态码
	Data any    `json:"data"` // 数据
	Msg  string `json:"msg"`  // 提示信息
}

// Ok 响应成功
func Ok(ctx iris.Context, data any) {
	ctx.JSON(iris.Map{"code": iris.StatusOK, "data": data, "msg": "ok"})
}

// Error 响应失败
func Error(ctx iris.Context, code int, data any, msg string) {
	ctx.JSON(iris.Map{"code": code, "data": data, "msg": msg})
}
