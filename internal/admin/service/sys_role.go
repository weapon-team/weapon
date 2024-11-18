package service

import (
	"github.com/weapon-team/weapon/internal/admin/helper"
	"github.com/weapon-team/weapon/internal/admin/model"
	"github.com/weapon-team/weapon/internal/sdk/base"
	"github.com/weapon-team/weapon/internal/sdk/engine"
	"github.com/weapon-team/weapon/internal/sdk/page"
)

// SysRoleService 系统角色逻辑 & 数据访问层
type SysRoleService struct {
	*engine.Engines
}

func NewSysRoleService(deps *engine.Engines) *SysRoleService {
	return &SysRoleService{deps}
}

// List 分页查询
func (s *SysRoleService) List(param helper.ListRoleParam) (page.Response, error) {
	var roles []model.SysRole
	session := s.Orm()
	if param.Name != "" {
		session.Where("name like?", "%"+param.Name+"%")
	}
	if param.Code != 0 {
		session.Where("code=?", param.Code)
	}
	count, err := session.Limit(param.Limit(), param.Offset()).FindAndCount(&roles)
	if err != nil {
		return page.Response{}, err
	}
	return page.NewResponse(param.Request, count, roles), nil
}

// Get 获取角色
func (s *SysRoleService) Get(id uint) (model.SysRole, error) {
	var role model.SysRole
	_, err := s.Orm().Where("id=?", id).Get(&role)
	return role, err
}

// Create 创建角色
func (s *SysRoleService) Create(param helper.CreateRoleParam, optUserId int64) (model.SysRole, error) {
	role := model.SysRole{
		Name:        param.Name,
		Code:        param.Code,
		DataScope:   param.DataScope,
		Description: param.Description,
		Sort:        param.Sort,
		Status:      param.Status,
		OptModel: base.OptModel{
			CreateUser: optUserId,
		},
	}
	_, err := s.Orm().InsertOne(&role)
	return role, err
}

// Update 修改用户
func (s *SysRoleService) Update(param helper.UpdateRoleParam, updateUser int64) (model.SysRole, error) {
	role := model.SysRole{
		Name:        param.Name,
		Code:        param.Code,
		DataScope:   param.DataScope,
		Description: param.Description,
		Sort:        param.Sort,
		Status:      param.Status,
		OptModel: base.OptModel{
			UpdateUser: updateUser,
		},
	}
	_, err := s.Orm().Update(&role)
	return role, err
}

func (s *SysRoleService) Delete(userId uint) error {
	_, err := s.Orm().Where("id=?", userId).Delete(&model.SysRole{})
	return err
}
