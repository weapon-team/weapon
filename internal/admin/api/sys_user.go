package api

import (
	"github.com/kataras/iris/v12"
	"github.com/redis/go-redis/v9"
	"xorm.io/xorm"

	"github.com/weapon-team/weapon/internal/admin/model"
	adminService "github.com/weapon-team/weapon/internal/admin/service"
	appService "github.com/weapon-team/weapon/internal/app/service"
	"github.com/weapon-team/weapon/internal/sdk/middleware/jwts"
	"github.com/weapon-team/weapon/internal/sdk/resp"
)

// SysUserApi 系统用户API层
type SysUserApi struct{}

// Hello 测试接口
// path: /admin/user/hello
func (e SysUserApi) Hello(ctx iris.Context, orm *xorm.Engine, rdb *redis.Client) {

	sysUserService := adminService.NewSysUserService(orm)
	appUserService := appService.NewAppUserService(orm)
	m := iris.Map{
		"SysUser": sysUserService.Hello(),
		"AppUser": appUserService.Hello(),
	}
	resp.Ok(ctx, m)
}

// Login 登录
// path: /admin/user/login
func (s SysUserApi) Login(ctx iris.Context, orm *xorm.Engine, rdb *redis.Client) {

	// TODO 接收参数
	us := adminService.NewSysUserService(orm)
	user := us.Login()
	token, err := jwts.GenerateToken(jwts.JwtClaims{User: user.Nickname})
	if err != nil {
		resp.Error(ctx, 1000, "", err.Error())
		return
	}
	resp.Ok(ctx, iris.Map{
		"user":  user,
		"token": token,
	})
}

// Create
// path: /admin/user/create
func (e SysUserApi) Create(ctx iris.Context, orm *xorm.Engine, rdb *redis.Client) {

	sysUserService := adminService.NewSysUserService(orm)
	su := model.SysUser{
		Username: "Im Jordan",
		Nickname: "JJ",
		Phone:    "110",
	}
	ok := sysUserService.Create(&su)
	if !ok {
		resp.Error(ctx, 1000, "", "create error")
		return
	}
	resp.Ok(ctx, su)
}
