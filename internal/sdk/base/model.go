package base

import (
	"time"
)

// BaseModel 基础模型
//
//	Xorm会自动处理创建时间、更新时间、删除时间: 创建时给定当前时间、更新时给定当前时间、删除时给定当前时间
type BaseModel struct {
	Created time.Time `xorm:"created"` // 创建时间
	Updated time.Time `xorm:"updated"` // 更新时间
	Deleted time.Time `xorm:"deleted"` // 删除时间
}
