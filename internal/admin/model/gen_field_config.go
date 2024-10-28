package model

import (
	"github.com/weapon-team/weapon/internal/sdk/base"
)

// GenFieldConfig 字段配置表
type GenFieldConfig struct {
	TableName      string `xorm:"not null comment('表名称') index VARCHAR(64)"`
	ColumnName     string `xorm:"not null comment('列名称') VARCHAR(64)"`
	ColumnType     string `xorm:"not null comment('列类型') VARCHAR(25)"`
	ColumnSize     int64  `xorm:"comment('列大小') BIGINT"`
	FieldName      string `xorm:"not null comment('字段名称') VARCHAR(64)"`
	FieldType      string `xorm:"not null comment('字段类型') VARCHAR(25)"`
	FieldSort      int    `xorm:"not null default 999 comment('字段排序') INT"`
	Comment        string `xorm:"comment('注释') VARCHAR(512)"`
	IsRequired     int    `xorm:"not null default b'1' comment('是否必填') BIT(1)"`
	ShowInList     int    `xorm:"not null default b'1' comment('是否在列表中显示') BIT(1)"`
	ShowInForm     int    `xorm:"not null default b'1' comment('是否在表单中显示') BIT(1)"`
	ShowInQuery    int    `xorm:"not null default b'1' comment('是否在查询中显示') BIT(1)"`
	FormType       uint   `xorm:"comment('表单类型') UNSIGNED TINYINT"`
	QueryType      uint   `xorm:"comment('查询方式') UNSIGNED TINYINT"`
	base.TimeModel `xorm:"extends"`
}
