package model

import "github.com/weapon-team/weapon/internal/sdk/base"

// SysOption 系统配置表
type SysOption struct {
	Name           string `xorm:"not null comment('名称') VARCHAR(50)"`
	Code           string `xorm:"not null pk comment('键') VARCHAR(100)"`
	Value          string `xorm:"comment('值') TEXT(65535)"`
	DefaultValue   string `xorm:"comment('默认值') TEXT(65535)"`
	Description    string `xorm:"comment('描述') VARCHAR(200)"`
	UpdateUser     int64  `xorm:"comment('修改人') BIGINT"`
	base.TimeModel `xorm:"extends"`
}
