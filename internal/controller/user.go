package controller

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "mall-api/api/v1"
	"mall-api/internal/service"
	"mall-api/utility"
)

var (
	User = cUser{}
)

type cUser struct{}

/**
登录
 */
func (c *cUser) Login(ctx context.Context, req *v1.UserLoginReq) (res *v1.UserLoginRes, err error) {
	if err = service.User().Login(ctx, req.LoginName, req.PasswordMd5); err != nil {
		return nil, gerror.New("login error!")
	}
	// 登录成功给客户端返回token
	res = &v1.UserLoginRes{}
	// 获取token
	res.Token, _ = utility.Auth().LoginHandler(ctx)
	return
}

/**
我的-用户信息
 */
func (c *cUser) GetUserInfo(ctx context.Context, req *v1.UserInfoReq) (*v1.UserInfoRes, error) {
	userId := gconv.String(utility.Auth().GetIdentity(ctx))
	return service.User().GetUserInfo(ctx, userId)
}