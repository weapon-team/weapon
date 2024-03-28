package service

import "xorm.io/xorm"

// App用户逻辑层
type AppUserService struct {
	orm *xorm.Engine
}

// New
func NewAppUserService(orm *xorm.Engine) *AppUserService {
	return &AppUserService{
		orm: orm,
	}
}
