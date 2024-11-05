package model

import (
	"github.com/weapon-team/weapon/internal/sdk/base"
)

// SysDept 部门表
type SysDept struct {
	base.AutoIncrId `xorm:"extends"`
	Name            string `xorm:"not null comment('名称') unique(uk_name_parent_id) VARCHAR(30)"`
	ParentId        int64  `xorm:"not null default 0 comment('上级部门ID') index unique(uk_name_parent_id) BIGINT"`
	Ancestors       string `xorm:"not null default '' comment('祖级列表') VARCHAR(512)"`
	Description     string `xorm:"comment('描述') VARCHAR(200)"`
	Sort            int    `xorm:"not null default 999 comment('排序') INT"`
	Status          uint   `xorm:"not null default 1 comment('状态（1：启用；2：禁用）') UNSIGNED TINYINT"`
	base.OptModel   `xorm:"extends"`
	base.TimeModel  `xorm:"extends"`
}

func (SysDept) TableName() string {
	return "sys_dept"
}
