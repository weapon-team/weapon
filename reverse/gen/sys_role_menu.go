package models

type SysRoleMenu struct {
	RoleId int64 `xorm:"not null pk comment('角色ID') BIGINT"`
	MenuId int64 `xorm:"not null pk comment('菜单ID') BIGINT"`
}
