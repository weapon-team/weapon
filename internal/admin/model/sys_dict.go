package model

import (
	"github.com/weapon-team/weapon/internal/sdk/base"
)

// SysDict 字典表
type SysDict struct {
	base.AutoIncrId `xorm:"extends"`
	Name            string        `xorm:"not null comment('名称') unique VARCHAR(30)"`
	Code            string        `xorm:"not null comment('编码') unique VARCHAR(30)"`
	Description     string        `xorm:"comment('描述') VARCHAR(200)"`
	Items           []SysDictItem `xorm:"not null default '{}' JSONB"`
	base.OptModel   `xorm:"extends"`
	base.TimeModel  `xorm:"extends"`
}

func (SysDict) TableName() string {
	return "sys_dict"
}

// SysDictItem 字典项表
type SysDictItem struct {
	Label string `xorm:"not null comment('标签') VARCHAR(30)"`
	Value string `xorm:"not null comment('值') unique(uk_value_dict_id) VARCHAR(30)"`
	Sort  int    `xorm:"not null default 999 comment('排序') INT"`
}
