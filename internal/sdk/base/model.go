package base

import "github.com/weapon-team/weapon/internal/sdk/types"

type (

	// AutoIncrId 自增id模型
	AutoIncrId struct {
		Id int64 `json:"id" xorm:"pk autoincr comment('ID') BIGINT"`
	}

	// OptModel 操作用户id模型
	OptModel struct {
		CreateUser int64 `json:"-" xorm:"comment('创建人') index BIGINT"`
		UpdateUser int64 `json:"-" xorm:"comment('修改人') index BIGINT"`
	}

	// TimeModel 基础模型
	// Xorm会自动处理创建时间、更新时间、删除时间: 创建时给定当前时间、更新时给定当前时间、删除时给定当前时间
	TimeModel struct {
		CreateTime types.Timestamp `json:"createTime" xorm:"created comment('创建时间')"` // 创建时间
		UpdateTime types.Timestamp `json:"-" xorm:"updated comment('修改时间')"`          // 更新时间
		DeleteTime types.Timestamp `json:"-" xorm:"deleted comment('删除时间')"`          // 删除时间
	}

	// ExtModel 通用扩展字段
	ExtModel struct {
		Ext types.Ext `json:"ext" xorm:"not null default '{}' JSONB"` // 扩展字段
	}
)
