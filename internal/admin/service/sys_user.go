package service

import (
	"time"

	"github.com/weapon-team/weapon/internal/admin/model"
	"github.com/weapon-team/weapon/internal/sdk/base"
	"github.com/weapon-team/weapon/internal/sdk/engine"
)

// SysUserService 系统用户逻辑 & 数据访问层
type SysUserService struct {
	egs *engine.Engines
}

func NewSysUserService(egs *engine.Engines) *SysUserService {
	return &SysUserService{egs: egs}
}

func (s *SysUserService) Hello() string {
	return "Hello SysUser!"
}

// Login 登录
func (s *SysUserService) Login() model.SysUser {

	return model.SysUser{
		Id:           1,
		Username:     "6666",
		Nickname:     "老六",
		Password:     "",
		Gender:       0,
		Email:        "",
		Phone:        "",
		Avatar:       "",
		Description:  "",
		Status:       0,
		IsSystem:     0,
		PwdResetTime: time.Now(),
		DeptId:       0,
		CreateUser:   0,
		UpdateUser:   0,
		BaseTimeModel: base.BaseTimeModel{
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
			DeleteTime: time.Now(),
		},
	}
}

func (s *SysUserService) Create(user *model.SysUser) bool {
	_, err := s.egs.Orm().InsertOne(user)
	return err == nil
}
