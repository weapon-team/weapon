package service

import "github.com/weapon-team/weapon/internal/sdk/engine"

// SysRoleService 系统角色逻辑 & 数据访问层
type SysRoleService struct {
	*engine.Engines
}

func NewSysRoleService(deps *engine.Engines) *SysRoleService {
	return &SysRoleService{deps}
}
