package service

import "github.com/weapon-team/weapon/internal/sdk/engine"

// SysRoleService 系统角色逻辑 & 数据访问层
type SysRoleService struct {
	ens *engine.Engines
}

func NewSysRoleService(ens *engine.Engines) *SysRoleService {
	return &SysRoleService{ens: ens}
}
