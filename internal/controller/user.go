package controller

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	v1 "mall-api/api/v1"
	"mall-api/internal/service"
)

var (
	User = cUser{}
)

type cUser struct{}

// 登录处理
func (c *cUser) Login(ctx context.Context, req *v1.UserLoginReq) (res *v1.UserLoginRes, err error) {
	if err = service.User().Login(ctx, req.LoginName, req.PasswordMd5); err != nil {
		return nil, gerror.New("login error!")
	}

	// 登录成功给客户端返回token
	res = &v1.UserLoginRes{}
	// 获取token
	res.Token, _ = service.Auth().LoginHandler(ctx)
	return
}
