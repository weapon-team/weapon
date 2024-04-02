package model

import (
	"github.com/weapon-team/weapon/internal/sdk/base"
)

// SysRole 角色表
type SysRole struct {
	Id             int64  `xorm:"pk autoincr comment('ID') BIGINT"`
	Name           string `xorm:"not null comment('名称') unique VARCHAR(30)"`
	Code           string `xorm:"not null comment('编码') unique VARCHAR(30)"`
	DataScope      int    `xorm:"not null default 4 comment('数据权限（1：全部数据权限；2：本部门及以下数据权限；3：本部门数据权限；4：仅本人数据权限；5：自定义数据权限）') TINYINT(1)"`
	Description    string `xorm:"comment('描述') VARCHAR(200)"`
	Sort           int    `xorm:"not null default 999 comment('排序') INT"`
	Status         uint   `xorm:"not null default 1 comment('状态（1：启用；2：禁用）') UNSIGNED TINYINT"`
	IsSystem       int    `xorm:"not null default b'0' comment('是否为系统内置数据') BIT(1)"`
	CreateUser     int64  `xorm:"not null comment('创建人') index BIGINT"`
	UpdateUser     int64  `xorm:"comment('修改人') index BIGINT"`
	base.BaseModel `xorm:"extends"`
}
