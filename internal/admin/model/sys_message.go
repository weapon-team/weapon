package model

import "github.com/weapon-team/weapon/internal/sdk/base"

// SysMessage 消息表
type SysMessage struct {
	Id                 int64  `xorm:"pk autoincr comment('ID') BIGINT"`
	Title              string `xorm:"not null comment('标题') VARCHAR(50)"`
	Content            string `xorm:"comment('内容') VARCHAR(255)"`
	Type               uint   `xorm:"not null default 1 comment('类型（1：系统消息）') UNSIGNED TINYINT"`
	CreateUser         int64  `xorm:"comment('创建人') BIGINT"`
	base.BaseTimeModel `xorm:"extends"`
}
