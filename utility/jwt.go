package utility

import (
	"context"
	jwt "github.com/gogf/gf-jwt/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "mall-api/api/v1"
	"mall-api/internal/service"
	"time"
)

var authService *jwt.GfJWTMiddleware

// 权限包管理
func Auth() *jwt.GfJWTMiddleware {
	return authService
}

// 初始化
func init() {
	auth := jwt.New(&jwt.GfJWTMiddleware{
		//用户的领域名称，必传
		Realm: "mall",
		// 签名算法
		SigningAlgorithm: "HS256",
		// 签名密钥
		Key: []byte("mall"),
		// 时效
		Timeout: time.Minute * 60,
		// 	token过期后，可凭借旧token获取新token的刷新时间
		MaxRefresh: time.Minute * 5,
		// 身份验证的key值
		IdentityKey: "userId",
		//token检索模式，用于提取token-> Authorization
		// TokenLookup: "header: Authorization, query: token, cookie: jwt",
		TokenLookup: "header: Authorization",
		// token在请求头时的名称，默认值为Bearer.客户端在header中传入"Authorization":"token xxxxxx"
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
		// 用户标识 map  私有属性
		// 根据登录信息对用户进行身份验证的回调函数
		Authenticator: Authenticator,
		// 处理不进行授权的逻辑
		Unauthorized: Unauthorized,
		//登录期间的设置私有载荷的函数，默认设置Authenticator函数回调的所有内容
		PayloadFunc: PayloadFunc,
		// 解析并设置用户身份信息，并设置身份信息至每次请求中
		IdentityHandler: IdentityHandler,
	})
	authService = auth
}

// 向webtoken添加额外的有效负载数据。
func PayloadFunc(data interface{}) jwt.MapClaims {
	claims := jwt.MapClaims{}
	// params := data.(map[string]interface{})
	params := gconv.Map(data)
	if len(params) > 0 {
		for k, v := range params {
			claims[k] = v
		}
	}
	return claims
}

// 标识
func IdentityHandler(ctx context.Context) interface{} {
	claims := jwt.ExtractClaims(ctx)
	return claims[authService.IdentityKey]
}

func Unauthorized(ctx context.Context, code int, message string) {
	r := g.RequestFromCtx(ctx)
	_ = r.Response.WriteJson(g.Map{
		"code":    code,
		"message": message,
	})
	r.ExitAll()
}

// 用户标识  私有载荷
func Authenticator(ctx context.Context) (interface{}, error) {
	var (
		r  = g.RequestFromCtx(ctx)
		in v1.UserLoginReq
	)
	if err := r.Parse(&in); err != nil {
		return "", err
	}

	if user := service.User().GetUserInfo(ctx, in.LoginName, in.PasswordMd5); user != nil {
		return user, nil
	}

	return nil, jwt.ErrFailedAuthentication
}


// 权限中间件
type middlewareService struct{}

var middleware = middlewareService{}

func Middleware() *middlewareService {
	return &middleware
}

func (s *middlewareService) Auth(r *ghttp.Request) {
	// GfJWTMiddleware gf jwt集成的中间件
	// Auth是权限service中配置的gf jwt
	Auth().MiddlewareFunc()(r)
	r.Middleware.Next()
}