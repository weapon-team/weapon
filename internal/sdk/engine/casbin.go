package engine

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	xormadapter "github.com/casbin/xorm-adapter/v2"
	"xorm.io/xorm"
)

const modelText = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = (r.sub == p.sub) && (keyMatch(r.obj, p.obj) || keyMatch2(r.obj, p.obj)) && (r.act == p.act || p.act == "*")
`

func initCasbinEnforcer(orm *xorm.Engine) (*casbin.SyncedEnforcer, error) {

	// 加载模型
	m, err := model.NewModelFromString(modelText)
	if err != nil {
		return nil, err
	}
	engine, err := xormadapter.NewAdapterByEngineWithTableName(orm, "sys_permission", "")
	if err != nil {
		return nil, err
	}
	enf, err := casbin.NewSyncedEnforcer(m, engine)
	if err != nil {
		return nil, err
	}
	err = enf.LoadPolicy()
	return enf, err
}
