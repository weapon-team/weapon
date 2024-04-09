package model

import "github.com/weapon-team/weapon/internal/sdk/base"

// SysMenu 菜单表
type SysMenu struct {
	Id                 int64  `xorm:"pk autoincr comment('ID') BIGINT"`
	Title              string `xorm:"not null comment('标题') unique(uk_title_parent_id) VARCHAR(30)"`
	ParentId           int64  `xorm:"not null default 0 comment('上级菜单ID') index unique(uk_title_parent_id) BIGINT"`
	Type               uint   `xorm:"not null default 1 comment('类型（1：目录；2：菜单；3：按钮）') UNSIGNED TINYINT"`
	Path               string `xorm:"comment('路由地址') VARCHAR(255)"`
	Name               string `xorm:"comment('组件名称') VARCHAR(50)"`
	Component          string `xorm:"comment('组件路径') VARCHAR(255)"`
	Icon               string `xorm:"comment('图标') VARCHAR(50)"`
	IsExternal         int    `xorm:"not null default b'0' comment('是否外链') BIT(1)"`
	IsCache            int    `xorm:"not null default b'0' comment('是否缓存') BIT(1)"`
	IsHidden           int    `xorm:"not null default b'0' comment('是否隐藏') BIT(1)"`
	Permission         string `xorm:"comment('权限标识') VARCHAR(100)"`
	Sort               int    `xorm:"not null default 999 comment('排序') INT"`
	Status             uint   `xorm:"not null default 1 comment('状态（1：启用；2：禁用）') UNSIGNED TINYINT"`
	CreateUser         int64  `xorm:"not null comment('创建人') index BIGINT"`
	UpdateUser         int64  `xorm:"comment('修改人') index BIGINT"`
	base.BaseTimeModel `xorm:"extends"`
}
