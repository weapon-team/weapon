package model

import "github.com/weapon-team/weapon/internal/sdk/base"

// 系统用户实体
type SysUser struct {
	Id             int64
	Name           string
	NickName       string
	Phone          string
	base.BaseModel `xorm:"extends"`
}
