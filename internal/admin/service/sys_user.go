package service

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/weapon-team/weapon/internal/admin/helper"
	"github.com/weapon-team/weapon/internal/admin/model"
	"github.com/weapon-team/weapon/internal/sdk/base"
	"github.com/weapon-team/weapon/internal/sdk/engine"
	"github.com/weapon-team/weapon/internal/sdk/page"
	"github.com/weapon-team/weapon/internal/sdk/types"
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
	//ok := s.Captcha().Verify(param.CaptchaId, param.Captcha, true)
	//if !ok {
	//	return sysUser, errors.New("验证码错误")
	//}
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

// List 分页查询
func (s *SysUserService) List(param helper.ListUserParam) (page.Response, error) {
	var users []model.SysUser
	session := s.Orm()
	if param.Username != "" {
		session.Where("username like?", "%"+param.Username+"%")
	}
	if param.Nickname != "" {
		session.Where("nickname like?", "%"+param.Nickname+"%")
	}
	if param.Gender != 0 {
		session.Where("gender=?", param.Gender)
	}
	count, err := session.Limit(param.Limit(), param.Offset()).FindAndCount(&users)
	if err != nil {
		return page.Response{}, err
	}
	return page.NewResponse(param.Request, count, users), nil
}

// Create 创建用户
func (s *SysUserService) Create(param helper.CreateUserParam) (model.SysUser, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(param.Password), bcrypt.DefaultCost)
	user := model.SysUser{
		Username:     param.Username,
		Nickname:     param.Nickname,
		Password:     string(password),
		Gender:       param.Gender,
		PwdResetTime: types.NowTimestamp(),
		OptModel: base.OptModel{
			CreateUser: 0,
			UpdateUser: 0,
		},
	}
	_, err = s.Orm().InsertOne(&user)
	return user, err
}

// Update 修改用户
func (s *SysUserService) Update(param helper.UpdateUserParam, updateUser int64) (model.SysUser, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(param.Password), bcrypt.DefaultCost)
	user := model.SysUser{
		Username: param.Username,
		Nickname: param.Nickname,
		Password: string(password),
		Gender:   param.Gender,
		//PwdResetTime: types.NowTimestamp(),
		OptModel: base.OptModel{
			UpdateUser: updateUser,
		},
	}
	_, err = s.Orm().Update(&user)
	return user, err
}

func (s *SysUserService) Delete(userId int64) error {
	_, err := s.Orm().Where("id=?", userId).Delete(&model.SysUser{})
	return err
}
