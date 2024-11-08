package middleware

import (
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"

	"github.com/weapon-team/weapon/internal/sdk/resp"
	"github.com/weapon-team/weapon/internal/sdk/runtime"
)

var signer *jwt.Signer // Jwt签名器

// InitJwt 初始化Jwt所需signer
func InitJwt() {
	signer = jwt.NewSigner(jwt.HS256, []byte(runtime.Setting.Jwt.Secret), time.Duration(runtime.Setting.Jwt.Expire)*time.Second)
}

// JwtClaims TODO 自定义claims
type JwtClaims struct {
	Uid      int64  `json:"uid"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

// GenerateToken 生成token
func GenerateToken(claims JwtClaims) (string, error) {
	token, err := signer.Sign(claims)
	if err != nil {
		return "", err
	}
	return string(token), nil
}

// JwtMiddleware Jwt鉴权中间件
func JwtMiddleware() iris.Handler {

	// Jwt验证器
	verifier := jwt.NewVerifier(jwt.HS256, []byte(runtime.Setting.Jwt.Secret))

	// 错误处理
	verifier.ErrorHandler = func(ctx iris.Context, err error) {
		resp.ErrorCtx(ctx, iris.StatusUnauthorized, "", err.Error()) // 返回401
		ctx.StopExecution()                                          // 停止后续执行
		//_ = ctx.StopWithJSON(iris.StatusOK, iris.Map{"code": iris.StatusUnauthorized, "data": "", "msg": "身份认证失败"})
	}
	return verifier.Verify(func() interface{} { return new(JwtClaims) })
}

// GetJwtClaims 获取Jwt中的claims
func GetJwtClaims(ctx iris.Context) *JwtClaims {
	return jwt.Get(ctx).(*JwtClaims)
}

// GetUserId 获取用户ID
func GetUserId(ctx iris.Context) int64 {
	claims := GetJwtClaims(ctx)
	return claims.Uid
}

// GetRemainTime 获取token剩余时间
func GetRemainTime(ctx iris.Context) time.Duration {
	scs := jwt.GetVerifiedToken(ctx).StandardClaims
	return scs.Timeleft()
}

// GetExpireTime 获取token过期时间
func GetExpireTime(ctx iris.Context) time.Time {
	scs := jwt.GetVerifiedToken(ctx).StandardClaims
	return scs.ExpiresAt()
}
