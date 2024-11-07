package internal

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/magiconair/properties/assert"
	"github.com/redis/go-redis/v9"
	"xorm.io/xorm"

	_ "github.com/go-sql-driver/mysql"

	"github.com/weapon-team/weapon/internal/admin/service"
	"github.com/weapon-team/weapon/internal/sdk/engine"
	"github.com/weapon-team/weapon/internal/sdk/xormlog"
)

func buildEngines() *engine.Engines {
	// 3. 数据库引擎
	orm, err := xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:8333)/continew_admin?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	orm.SetLogger(&xormlog.XormLogger{})
	orm.ShowSQL(true)
	if err := orm.Ping(); err != nil {
		panic(err)
	}
	// 4. redis缓存
	rdb := redis.NewClient(&redis.Options{
		Addr:           "120.55.53.30:7600",
		Username:       "",
		Password:       "",
		MaxActiveConns: 20,
	})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		panic(fmt.Errorf("failed to connect to redis: %w", err))
	}

	// 5. casbin
	//casbinEngine, err := xormadapter.NewAdapterByEngine(orm)
	//if err != nil {
	//	panic(err)
	//}
	//enf, err := casbin.NewSyncedEnforcer("rbac_models.conf", casbinEngine)
	//if err != nil {
	//	panic(err)
	//}
	//err = enf.LoadPolicy()
	//if err != nil {
	//	panic(err)
	//}

	return engine.NewEngines(orm, rdb, nil)
}

func TestService(t *testing.T) {

	ens := buildEngines()

	optSve := service.NewSysOptionService(ens)
	options, err := optSve.AllOptions()
	assert.Equal(t, err, nil, "")

	for i, v := range options {
		t.Logf("[%v]Option: %v", i, v)
	}

}

type User struct {
	Username string `json:"username" validate:"required" msg:"用户名不能为空"`
	Password string `json:"password" validate:"required"`
}

func TestValidate(t *testing.T) {

	u := User{
		Username: "",
		Password: "",
	}
	valgo := validator.New()
	err := valgo.RegisterValidation("msg", func(fl validator.FieldLevel) bool {
		return fl.Field().Equal(reflect.ValueOf("username"))
	}, true)
	if err != nil {
		t.Fatal("20000-------- ", err)
	}
	err = valgo.Struct(&u)
	t.Log(err)
}
