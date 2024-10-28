package model

import (
	"github.com/weapon-team/weapon/internal/sdk/base"
	"github.com/weapon-team/weapon/internal/sdk/types"
)

// SysUser 系统用户表
type SysUser struct {
	base.AutoIncrId `xorm:"extends"`
	Username        string          `xorm:"not null comment('用户名') unique VARCHAR(64)"`
	Nickname        string          `xorm:"not null comment('昵称') VARCHAR(30)"`
	Password        string          `xorm:"comment('密码') VARCHAR(255)"`
	Gender          uint            `xorm:"not null default 0 comment('性别（0：未知；1：男；2：女）') UNSIGNED TINYINT"`
	Email           string          `xorm:"comment('邮箱') unique VARCHAR(255)"`
	Phone           string          `xorm:"comment('手机号码') unique VARCHAR(255)"`
	Avatar          string          `xorm:"comment('头像地址') VARCHAR(255)"`
	Description     string          `xorm:"comment('描述') VARCHAR(200)"`
	Status          uint            `xorm:"not null default 1 comment('状态（1：启用；2：禁用）') UNSIGNED TINYINT"`
	PwdResetTime    types.Timestamp `xorm:"comment('最后一次修改密码时间')"`
	DeptId          int64           `xorm:"not null comment('部门ID') index BIGINT"`
	base.OptModel   `xorm:"extends"`
	base.TimeModel  `xorm:"extends"`
	base.ExtModel   `xorm:"extends"`
}
