package model

import (
	"github.com/weapon-team/weapon/internal/sdk/base"
	"github.com/weapon-team/weapon/internal/sdk/types"
)

// SysAnnouncement 公告表
type SysAnnouncement struct {
	base.AutoIncrId `xorm:"extends"`
	Title           string          `xorm:"not null comment('标题') VARCHAR(150)"`
	Content         string          `xorm:"not null comment('内容') MEDIUMTEXT(16777215)"`
	Type            string          `xorm:"not null comment('类型') VARCHAR(30)"`
	EffectiveTime   types.Timestamp `xorm:"comment('生效时间')"`
	TerminateTime   types.Timestamp `xorm:"comment('终止时间')"`
	Sort            int             `xorm:"not null default 999 comment('排序') INT"`
	base.OptModel   `xorm:"extends"`
	base.TimeModel  `xorm:"extends"`
}
