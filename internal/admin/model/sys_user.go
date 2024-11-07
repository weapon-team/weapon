package model

import (
	"github.com/weapon-team/weapon/internal/sdk/base"
	"github.com/weapon-team/weapon/internal/sdk/types"
)

// SysUser 系统用户表
type SysUser struct {
	base.AutoIncrId `xorm:"extends"`
	Username        string          `json:"username" xorm:"not null comment('用户名') unique VARCHAR(64)"`
	Nickname        string          `json:"nickname" xorm:"not null comment('昵称') VARCHAR(30)"`
	Password        string          `json:"-" xorm:"comment('密码') VARCHAR(255)"`
	Gender          uint            `json:"gender" xorm:"not null default 0 comment('性别（0：未知；1:男；2: 女）') UNSIGNED TINYINT"`
	Email           string          `json:"email" xorm:"comment('邮箱') unique VARCHAR(255)"`
	Phone           string          `json:"phone" xorm:"comment('手机号码') unique VARCHAR(255)"`
	Avatar          string          `json:"avatar" xorm:"comment('头像地址') VARCHAR(255)"`
	Status          uint            `json:"status" xorm:"not null default 1 comment('状态（1：启用；2：禁用）') UNSIGNED TINYINT"`
	PwdResetTime    types.Timestamp `json:"pwdResetTime" xorm:"comment('最后一次修改密码时间')"`
	DeptId          int64           `json:"deptId" xorm:"not null comment('部门ID') index BIGINT"`
	base.OptModel   `xorm:"extends"`
	base.TimeModel  `xorm:"extends"`
	base.ExtModel   `xorm:"extends"`
}

func (SysUser) TableName() string {
	return "sys_user"
}
