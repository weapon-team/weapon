package xormlog

import (
	"fmt"

	"xorm.io/xorm/log"

	"github.com/weapon-team/weapon/pkg/logs"
)

// XormLogger
// 实现xorm的Logger接口,调用自行封装的logs
type XormLogger struct {
	level   log.LogLevel
	showSQL bool
}

var _ log.Logger = &XormLogger{}

func (x *XormLogger) Debug(v ...interface{}) {
	logs.Debug().Str("Type", "Xorm").Msg(fmt.Sprintln(v...))
}

func (x *XormLogger) Debugf(format string, v ...interface{}) {
	logs.Debug().Str("Type", "Xorm").Msgf(fmt.Sprintf(format, v...))
}

func (x *XormLogger) Error(v ...interface{}) {
	logs.Error().Str("Type", "Xorm").Msg(fmt.Sprintln(v...))
}

func (x *XormLogger) Errorf(format string, v ...interface{}) {
	logs.Error().Str("Type", "Xorm").Msgf(fmt.Sprintf(format, v...))
}

func (x *XormLogger) Info(v ...interface{}) {
	logs.Info().Str("Type", "Xorm").Msg(fmt.Sprintln(v...))
}

func (x *XormLogger) Infof(format string, v ...interface{}) {
	logs.Info().Str("Type", "Xorm").Msgf(fmt.Sprintf(format, v...))
}

func (x *XormLogger) Warn(v ...interface{}) {
	logs.Warn().Str("Type", "Xorm").Msg(fmt.Sprintln(v...))
}

func (x *XormLogger) Warnf(format string, v ...interface{}) {
	logs.Warn().Str("Type", "Xorm").Msgf(fmt.Sprintf(format, v...))
}

func (x *XormLogger) Level() log.LogLevel {
	return x.level
}

func (x *XormLogger) SetLevel(l log.LogLevel) {
	x.level = l
}

func (x *XormLogger) ShowSQL(show ...bool) {
	if len(show) == 0 {
		x.showSQL = true
		return
	}
	x.showSQL = show[0]
}

func (x *XormLogger) IsShowSQL() bool {
	return x.showSQL
}
