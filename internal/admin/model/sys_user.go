package model

import (
	"time"

	"github.com/weapon-team/weapon/internal/sdk/base"
)

// SysUser 系统用户表
type SysUser struct {
	Id             int64          `xorm:"pk autoincr comment('ID') BIGINT"`
	Username       string         `xorm:"not null comment('用户名') unique VARCHAR(64)"`
	Nickname       string         `xorm:"not null comment('昵称') VARCHAR(30)"`
	Password       string         `xorm:"comment('密码') VARCHAR(255)"`
	Gender         uint           `xorm:"not null default 0 comment('性别（0：未知；1：男；2：女）') UNSIGNED TINYINT"`
	Email          string         `xorm:"comment('邮箱') unique VARCHAR(255)"`
	Phone          string         `xorm:"comment('手机号码') unique VARCHAR(255)"`
	Avatar         string         `xorm:"comment('头像地址') VARCHAR(255)"`
	Description    string         `xorm:"comment('描述') VARCHAR(200)"`
	Status         uint           `xorm:"not null default 1 comment('状态（1：启用；2：禁用）') UNSIGNED TINYINT"`
	PwdResetTime   time.Time      `xorm:"comment('最后一次修改密码时间') TIMESTAMPZ"`
	DeptId         int64          `xorm:"not null comment('部门ID') index BIGINT"`
	Ext            map[string]any `xorm:"default '{}' JSONB"`
	CreateUser     int64          `xorm:"comment('创建人') index BIGINT"`
	UpdateUser     int64          `xorm:"comment('修改人') index BIGINT"`
	base.TimeModel `xorm:"extends"`
}
