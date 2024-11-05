package model

import "github.com/weapon-team/weapon/internal/sdk/base"

// SysOption 系统配置表
type SysOption struct {
	Code           string `xorm:"not null pk comment('键') VARCHAR(100)"`
	Name           string `xorm:"not null comment('名称') VARCHAR(50)"`
	Value          string `xorm:"comment('值')"`
	DefaultValue   string `xorm:"comment('默认值')"`
	Description    string `xorm:"comment('描述') VARCHAR(200)"`
	base.OptModel  `xorm:"extends"`
	base.TimeModel `xorm:"extends"`
}

func (SysOption) TableName() string {
	return "sys_option"
}
