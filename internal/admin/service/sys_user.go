package service

import (
	"time"

	"github.com/weapon-team/weapon/internal/admin/model"
	"github.com/weapon-team/weapon/internal/sdk/base"
	"github.com/weapon-team/weapon/internal/sdk/engine"
)

// SysUserService 系统用户逻辑 & 数据访问层
type SysUserService struct {
	ens *engine.Engines
}

func NewSysUserService(ens *engine.Engines) *SysUserService {
	return &SysUserService{ens: ens}
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
		PwdResetTime: time.Now(),
		DeptId:       0,
		CreateUser:   0,
		UpdateUser:   0,
		TimeModel: base.TimeModel{
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
			DeleteTime: time.Now(),
		},
	}
}

func (s *SysUserService) Create(user *model.SysUser) bool {
	_, err := s.ens.Orm().InsertOne(user)
	return err == nil
}
