package models

import (
	"time"
)

type SysMessage struct {
	Id         int64     `xorm:"pk comment('ID') BIGINT"`
	Title      string    `xorm:"not null comment('标题') VARCHAR(50)"`
	Content    string    `xorm:"comment('内容') VARCHAR(255)"`
	Type       uint      `xorm:"not null default 1 comment('类型（1：系统消息）') UNSIGNED TINYINT"`
	CreateUser int64     `xorm:"comment('创建人') BIGINT"`
	CreateTime time.Time `xorm:"not null comment('创建时间') DATETIME"`
}
