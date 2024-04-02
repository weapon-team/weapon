package models

import (
	"time"
)

type SysAnnouncement struct {
	Id            int64     `xorm:"pk comment('ID') BIGINT"`
	Title         string    `xorm:"not null comment('标题') VARCHAR(150)"`
	Content       string    `xorm:"not null comment('内容') MEDIUMTEXT(16777215)"`
	Type          string    `xorm:"not null comment('类型') VARCHAR(30)"`
	EffectiveTime time.Time `xorm:"comment('生效时间') DATETIME"`
	TerminateTime time.Time `xorm:"comment('终止时间') DATETIME"`
	Sort          int       `xorm:"not null default 999 comment('排序') INT"`
	CreateUser    int64     `xorm:"not null comment('创建人') index BIGINT"`
	CreateTime    time.Time `xorm:"not null comment('创建时间') DATETIME"`
	UpdateUser    int64     `xorm:"comment('修改人') index BIGINT"`
	UpdateTime    time.Time `xorm:"comment('修改时间') DATETIME"`
}
