package model

import (
	"github.com/weapon-team/weapon/internal/sdk/base"
)

// SysDictItem 字典项表
type SysDictItem struct {
	Id             int64  `xorm:"pk autoincr comment('ID') BIGINT"`
	Label          string `xorm:"not null comment('标签') VARCHAR(30)"`
	Value          string `xorm:"not null comment('值') unique(uk_value_dict_id) VARCHAR(30)"`
	Color          string `xorm:"comment('标签颜色') VARCHAR(30)"`
	Sort           int    `xorm:"not null default 999 comment('排序') INT"`
	Description    string `xorm:"comment('描述') VARCHAR(200)"`
	DictId         int64  `xorm:"not null comment('字典ID') index unique(uk_value_dict_id) BIGINT"`
	CreateUser     int64  `xorm:"not null comment('创建人') index BIGINT"`
	UpdateUser     int64  `xorm:"comment('修改人') index BIGINT"`
	base.BaseModel `xorm:"extends"`
}
