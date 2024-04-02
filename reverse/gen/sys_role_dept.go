package models

type SysRoleDept struct {
	RoleId int64 `xorm:"not null pk comment('角色ID') BIGINT"`
	DeptId int64 `xorm:"not null pk comment('部门ID') BIGINT"`
}
