package model

import (
	"github.com/weapon-team/weapon/internal/sdk/base"
)

// SysDict 字典表
type SysDict struct {
	Id                 int64  `xorm:"pk autoincr comment('ID') BIGINT"`
	Name               string `xorm:"not null comment('名称') unique VARCHAR(30)"`
	Code               string `xorm:"not null comment('编码') unique VARCHAR(30)"`
	Description        string `xorm:"comment('描述') VARCHAR(200)"`
	IsSystem           int    `xorm:"not null default b'0' comment('是否为系统内置数据') BIT(1)"`
	CreateUser         int64  `xorm:"not null comment('创建人') BIGINT"`
	UpdateUser         int64  `xorm:"comment('修改人') BIGINT"`
	base.BaseTimeModel `xorm:"extends"`
}
