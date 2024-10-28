package model

import (
	"github.com/weapon-team/weapon/internal/sdk/base"
)

// SysDict 字典表
type SysDict struct {
	base.AutoIncrId `xorm:"extends"`
	Name            string `xorm:"not null comment('名称') unique VARCHAR(30)"`
	Code            string `xorm:"not null comment('编码') unique VARCHAR(30)"`
	Description     string `xorm:"comment('描述') VARCHAR(200)"`
	base.OptModel   `xorm:"extends"`
	base.TimeModel  `xorm:"extends"`
}
