package service

import (
	"github.com/weapon-team/weapon/internal/admin/model"
	"xorm.io/xorm"
)

// 系统用户逻辑层
type SysUserService struct {
	orm *xorm.Engine
}

func NewSysUserService(orm *xorm.Engine) *SysUserService {
	return &SysUserService{
		orm: orm,
	}
}

func (s *SysUserService) Hello() string {
	return "Hi, girl !"
}

func (s *SysUserService) Create(user *model.SysUser) bool {
	_, err := s.orm.InsertOne(user)
	if err != nil {
		return false
	}
	return true
}
