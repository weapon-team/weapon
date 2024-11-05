package model

// SysUserRole 用户角色关联表
type SysUserRole struct {
	UserId int64 `xorm:"not null pk comment('用户ID') BIGINT"`
	RoleId int64 `xorm:"not null pk comment('角色ID') BIGINT"`
}

func (SysUserRole) TableName() string {
	return "sys_user_role"
}
