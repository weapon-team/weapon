package middleware

import (
	"time"

	"github.com/kataras/iris/v12/middleware/logger"

	"github.com/weapon-team/weapon/pkg/logs"
)

// MyLogConfig 自定义log配置
func MyLogConfig() logger.Config {
	return logger.Config{
		Status:             true,
		IP:                 true,
		Method:             true,
		Path:               true,
		PathAfterHandler:   false,
		Query:              true,
		TraceRoute:         true,
		MessageContextKeys: nil,
		MessageHeaderKeys:  nil,
		LogFunc: func(endTime time.Time, latency time.Duration, status, ip, method, path string, message interface{}, headerMessage interface{}) {
			logs.Info().Msgf("[Iris] Request [%v][%v] %v %v | %v, message: %v, headerMessage: %v", method, status, ip, path, latency, message, headerMessage)
		},
		LogFuncCtx: nil,
		Skippers:   nil,
	}
}
