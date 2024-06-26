package model

import (
	"time"
)

// SysMessageUser 消息用户关联表
type SysMessageUser struct {
	MessageId int64     `xorm:"not null pk comment('消息ID') BIGINT"`
	UserId    int64     `xorm:"not null pk comment('用户ID') BIGINT"`
	IsRead    int       `xorm:"not null default b'0' comment('是否已读') BIT(1)"`
	ReadTime  time.Time `xorm:"comment('读取时间') DATETIME"`
}
