package resp

import (
	"time"

	"github.com/kataras/iris/v12"
)

// Resp 响应体结构
type Resp struct {
	Code      int    `json:"code"` // 状态码
	Data      any    `json:"data"` // 数据
	Msg       string `json:"msg"`  // 提示信息
	Timestamp int64  `json:"timestamp"`
}

// OkCtx 响应成功
func OkCtx(ctx iris.Context, data any) {
	_ = ctx.JSON(iris.Map{"code": iris.StatusOK, "data": data, "msg": "ok"})
}

// ErrorCtx 响应失败
func ErrorCtx(ctx iris.Context, code int, data any, msg string) {
	_ = ctx.JSON(iris.Map{"code": code, "data": data, "msg": msg})
}

// Ok 响应成功
func Ok(data any) Resp {
	return Resp{
		Code:      iris.StatusOK,
		Data:      data,
		Msg:       "ok",
		Timestamp: time.Now().Unix(),
	}
}

// Error 响应失败
func Error(code int, msg string) Resp {
	return Resp{
		Code:      code,
		Data:      iris.Map{},
		Msg:       msg,
		Timestamp: time.Now().Unix(),
	}
}
