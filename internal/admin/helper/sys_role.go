package helper

import "github.com/weapon-team/weapon/internal/sdk/page"

// ListRoleParam 分页查询角色请求参数
type ListRoleParam struct {
	page.Request
	Name string `json:"name"`
	Code int    `json:"code"`
}

// CreateRoleParam 创建角色请求参数
type CreateRoleParam struct {
	Name        string `form:"name" json:"name" validate:"required" error:"请输入角色名"`
	Code        string `form:"code" json:"code" validate:"required" error:"请输入角色码"`
	DataScope   uint   `form:"dataScope" json:"dataScope" validate:"oneof=1 2 3 4 5" error:"请选择数据权限"`
	Description string `form:"description" json:"description"`
	Sort        int    `form:"sort" json:"sort"`
	Status      uint   `form:"status" json:"status" validate:"oneof=1 2" error:"请选择状态"`
}

// UpdateRoleParam 修改角色请求参数
type UpdateRoleParam struct {
	Id          int64  `form:"id" json:"id" validate:"required" error:"invalid role id"`
	Name        string `form:"name" json:"name" validate:"required" error:"请输入角色名"`
	Code        string `form:"code" json:"code" validate:"required" error:"请输入角色码"`
	DataScope   uint   `form:"dataScope" json:"dataScope" validate:"oneof=1 2 3 4 5" error:"请选择数据权限"`
	Description string `form:"description" json:"description"`
	Sort        int    `form:"sort" json:"sort"`
	Status      uint   `form:"status" json:"status" validate:"oneof=1 2" error:"请选择状态"`
}
