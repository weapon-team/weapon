package service

import (
	"xorm.io/xorm"

	"github.com/weapon-team/weapon/internal/admin/model"
)

// SysUserService 系统用户逻辑 & 数据访问层
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
	return err == nil
}
