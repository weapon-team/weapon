package models

import (
	"time"
)

type SysUser struct {
	Id           int64     `xorm:"pk comment('ID') BIGINT"`
	Username     string    `xorm:"not null comment('用户名') unique VARCHAR(64)"`
	Nickname     string    `xorm:"not null comment('昵称') VARCHAR(30)"`
	Password     string    `xorm:"comment('密码') VARCHAR(255)"`
	Gender       uint      `xorm:"not null default 0 comment('性别（0：未知；1：男；2：女）') UNSIGNED TINYINT"`
	Email        string    `xorm:"comment('邮箱') unique VARCHAR(255)"`
	Phone        string    `xorm:"comment('手机号码') unique VARCHAR(255)"`
	Avatar       string    `xorm:"comment('头像地址') VARCHAR(255)"`
	Description  string    `xorm:"comment('描述') VARCHAR(200)"`
	Status       uint      `xorm:"not null default 1 comment('状态（1：启用；2：禁用）') UNSIGNED TINYINT"`
	IsSystem     int       `xorm:"not null default b'0' comment('是否为系统内置数据') BIT(1)"`
	PwdResetTime time.Time `xorm:"comment('最后一次修改密码时间') DATETIME"`
	DeptId       int64     `xorm:"not null comment('部门ID') index BIGINT"`
	CreateUser   int64     `xorm:"comment('创建人') index BIGINT"`
	CreateTime   time.Time `xorm:"not null comment('创建时间') DATETIME"`
	UpdateUser   int64     `xorm:"comment('修改人') index BIGINT"`
	UpdateTime   time.Time `xorm:"comment('修改时间') DATETIME"`
}
