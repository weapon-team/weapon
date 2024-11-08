package service

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/weapon-team/weapon/internal/admin/helper"
	"github.com/weapon-team/weapon/internal/admin/model"
	"github.com/weapon-team/weapon/internal/sdk/engine"
)

// SysUserService 系统用户逻辑 & 数据访问层
type SysUserService struct {
	*engine.Engines
}

func NewSysUserService(ens *engine.Engines) *SysUserService {
	return &SysUserService{ens}
}

func (s *SysUserService) Hello() string {
	return "Hello SysUser!"
}

// Login 登录
func (s *SysUserService) Login(param helper.LoginParam) (model.SysUser, error) {
	var sysUser model.SysUser
	ok := s.Captcha().Verify(param.CaptchaId, param.Captcha, true)
	if !ok {
		return sysUser, errors.New("验证码错误")
	}
	// 查询用户
	ok, err := s.Orm().Where("username=?", param.Username).Get(&sysUser)
	if err != nil || !ok {
		return sysUser, errors.New("用户不存在")
	}
	// 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(sysUser.Password), []byte(param.Password))
	if err != nil {
		return sysUser, errors.New("密码错误")
	}
	return sysUser, nil
}

func (s *SysUserService) Create(user *model.SysUser) bool {
	_, err := s.Orm().InsertOne(user)
	return err == nil
}
