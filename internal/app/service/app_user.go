package service

import (
	"github.com/weapon-team/weapon/internal/sdk/engine"
)

// AppUserService App用户逻辑层
type AppUserService struct {
	egs *engine.Engines
}

// NewAppUserService 新建App用户逻辑层
func NewAppUserService(egs *engine.Engines) *AppUserService {
	return &AppUserService{egs: egs}
}

func (s *AppUserService) Hello() string {

	return "Hi, AppUser !"
}
