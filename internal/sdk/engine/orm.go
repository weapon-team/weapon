package engine

import (
	"xorm.io/xorm"

	"github.com/weapon-team/weapon/internal/sdk/runtime"
	"github.com/weapon-team/weapon/internal/sdk/xormlog"
)

func initOrm() (*xorm.Engine, error) {
	orm, err := xorm.NewEngine(runtime.Setting.DataSource.Driver, runtime.Setting.DataSource.DSN)
	if err != nil {
		return nil, err
	}
	orm.SetLogger(&xormlog.XormLogger{})
	orm.ShowSQL(runtime.Setting.Log.ShowSQL)
	if err = orm.Ping(); err != nil {
		return nil, err
	}
	return orm, nil
}
