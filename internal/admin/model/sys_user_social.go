package model

import (
	"time"
)

// SysUserSocial 用户社交关联表
type SysUserSocial struct {
	Source        string    `xorm:"not null pk comment('来源') unique(uk_source_open_id) VARCHAR(255)"`
	OpenId        string    `xorm:"not null pk comment('开放ID') unique(uk_source_open_id) VARCHAR(255)"`
	UserId        int64     `xorm:"not null comment('用户ID') BIGINT"`
	MetaJson      string    `xorm:"comment('附加信息') TEXT(65535)"`
	LastLoginTime time.Time `xorm:"comment('最后登录时间') DATETIME"`
	CreateTime    time.Time `xorm:"not null comment('创建时间') DATETIME"`
}
