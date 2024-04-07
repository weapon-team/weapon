package service

import (
	"github.com/weapon-team/weapon/internal/sdk/dep"
)

// AppUserService App用户逻辑层
type AppUserService struct {
	deps *dep.Dependency
}

// NewAppUserService 新建App用户逻辑层
func NewAppUserService(deps *dep.Dependency) *AppUserService {
	return &AppUserService{deps: deps}
}

func (s *AppUserService) Hello() string {

	return "Hi, AppUser !"
}
