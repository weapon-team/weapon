package model

// SysRoleMenu 角色菜单关联表
type SysRoleMenu struct {
	RoleId int64 `xorm:"not null pk comment('角色ID') BIGINT"`
	MenuId int64 `xorm:"not null pk comment('菜单ID') BIGINT"`
}

func (SysRoleMenu) TableName() string {
	return "sys_role_menu"
}
