package model

import "github.com/weapon-team/weapon/internal/sdk/base"

// SysMenu 菜单表
type SysMenu struct {
	base.AutoIncrId `xorm:"extends"`
	Title           string `xorm:"not null comment('标题') unique(uk_title_parent_id) VARCHAR(30)"`
	ParentId        int64  `xorm:"not null default 0 comment('上级菜单ID') index unique(uk_title_parent_id) BIGINT"`
	Type            uint   `xorm:"not null default 1 comment('类型（1：目录；2：菜单；3：按钮）') UNSIGNED TINYINT"`
	Path            string `xorm:"comment('路由地址') VARCHAR(255)"`
	Name            string `xorm:"comment('组件名称') VARCHAR(50)"`
	Component       string `xorm:"comment('组件路径') VARCHAR(255)"`
	Icon            string `xorm:"comment('图标') VARCHAR(50)"`
	Permission      string `xorm:"comment('权限标识') VARCHAR(100)"`
	Sort            int    `xorm:"not null default 999 comment('排序') INT"`
	Status          uint   `xorm:"not null default 1 comment('状态（1：启用；2：禁用）') UNSIGNED TINYINT"`
	base.OptModel   `xorm:"extends"`
	base.TimeModel  `xorm:"extends"`
}

func (SysMenu) TableName() string {
	return "sys_menu"
}
