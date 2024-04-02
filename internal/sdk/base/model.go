package base

import (
	"time"
)

// BaseModel 基础模型
//
//	Xorm会自动处理创建时间、更新时间、删除时间: 创建时给定当前时间、更新时给定当前时间、删除时给定当前时间
type BaseModel struct {
	CreateTime time.Time `xorm:"created comment('创建时间')"` // 创建时间
	UpdateTime time.Time `xorm:"updated comment('修改时间')"` // 更新时间
	DeleteTime time.Time `xorm:"deleted comment('删除时间')"` // 删除时间
}
