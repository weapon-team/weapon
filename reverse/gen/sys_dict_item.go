package models

import (
	"time"
)

type SysDictItem struct {
	Id          int64     `xorm:"pk comment('ID') BIGINT"`
	Label       string    `xorm:"not null comment('标签') VARCHAR(30)"`
	Value       string    `xorm:"not null comment('值') unique(uk_value_dict_id) VARCHAR(30)"`
	Color       string    `xorm:"comment('标签颜色') VARCHAR(30)"`
	Sort        int       `xorm:"not null default 999 comment('排序') INT"`
	Description string    `xorm:"comment('描述') VARCHAR(200)"`
	DictId      int64     `xorm:"not null comment('字典ID') index unique(uk_value_dict_id) BIGINT"`
	CreateUser  int64     `xorm:"not null comment('创建人') index BIGINT"`
	CreateTime  time.Time `xorm:"not null comment('创建时间') DATETIME"`
	UpdateUser  int64     `xorm:"comment('修改人') index BIGINT"`
	UpdateTime  time.Time `xorm:"comment('修改时间') DATETIME"`
}
