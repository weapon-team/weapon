package model

import (
	"github.com/weapon-team/weapon/internal/sdk/base"
)

// GenConfig 生成配置表
type GenConfig struct {
	TableName          string `xorm:"not null pk comment('表名称') VARCHAR(64)"`
	ModuleName         string `xorm:"not null comment('模块名称') VARCHAR(60)"`
	PackageName        string `xorm:"not null comment('包名称') VARCHAR(60)"`
	BusinessName       string `xorm:"not null comment('业务名称') VARCHAR(50)"`
	Author             string `xorm:"not null comment('作者') VARCHAR(100)"`
	TablePrefix        string `xorm:"comment('表前缀') VARCHAR(20)"`
	IsOverride         int    `xorm:"not null default b'0' comment('是否覆盖') BIT(1)"`
	base.BaseTimeModel `xorm:"extends"`
}
