package service

import (
	"github.com/weapon-team/weapon/internal/admin/helper"
	"github.com/weapon-team/weapon/internal/admin/model"
	"github.com/weapon-team/weapon/internal/sdk/engine"
	"github.com/weapon-team/weapon/pkg/ternary"
)

// SysOptionService 系统配置逻辑 & 数据访问层
type SysOptionService struct {
	egs *engine.Engines
}

func NewSysOptionService(egs *engine.Engines) *SysOptionService {
	return &SysOptionService{egs: egs}
}

// AllOptions 获取所有配置
func (s *SysOptionService) AllOptions() ([]helper.SysOptionResp, error) {

	// 查询
	var options []model.SysOption
	if err := s.egs.Orm().Find(&options); err != nil {
		return nil, err
	}

	// 组装
	result := make([]helper.SysOptionResp, 0, len(options))
	for _, option := range options {
		sor := helper.SysOptionResp{
			Label:    option.Code,
			Value:    ternary.If(option.Value != "", option.Value, option.DefaultValue),
			Disabled: true,
		}
		result = append(result, sor)
	}
	return result, nil
}
