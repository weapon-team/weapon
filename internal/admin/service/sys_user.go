package service

import (
	"time"

	"xorm.io/xorm"

	"github.com/weapon-team/weapon/internal/admin/model"
	"github.com/weapon-team/weapon/internal/sdk/base"
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
	return "Hi, SysUser !"
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
		BaseModel: base.BaseModel{
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
			DeleteTime: time.Now(),
		},
	}
}

func (s *SysUserService) Create(user *model.SysUser) bool {
	_, err := s.orm.InsertOne(user)
	return err == nil
}
