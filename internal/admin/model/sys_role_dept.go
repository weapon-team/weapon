package model

// SysRoleDept 角色部门关联表
type SysRoleDept struct {
	RoleId int64 `xorm:"not null pk comment('角色ID') BIGINT"`
	DeptId int64 `xorm:"not null pk comment('部门ID') BIGINT"`
}

func (SysRoleDept) TableName() string {
	return "sys_role_dept"
}
