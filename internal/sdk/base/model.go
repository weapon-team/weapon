package base

import (
	"time"
)

type BaseModel struct {
	Created time.Time `xorm:"created"` // 创建时间
	Updated time.Time `xorm:"updated"` // 更新时间
	Deleted time.Time `xorm:"deleted"` // 删除时间
}
