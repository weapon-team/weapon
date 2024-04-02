package service

import "xorm.io/xorm"

// AppUserService App用户逻辑层
type AppUserService struct {
	orm *xorm.Engine
}

// NewAppUserService 新建App用户逻辑层
func NewAppUserService(orm *xorm.Engine) *AppUserService {
	return &AppUserService{
		orm: orm,
	}
}
