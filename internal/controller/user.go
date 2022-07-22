package controller

import (
	"context"
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
func (c *cUser) Login(ctx context.Context, req *v1.UserLoginReq) (*v1.UserLoginRes, error) {
	var token *v1.UserLoginRes
	if err := service.User().Login(ctx, req.LoginName, req.PasswordMd5); err != nil {
		return nil, err
	}
	//tokenString, _:= utility.Auth().LoginHandler(ctx)
	//err := gconv.Scan(tokenString, &token)
	//if err!=nil{
	//	return nil, err
	//}
	return token, nil
}

/**
我的-用户信息
*/
func (c *cUser) GetUserInfo(ctx context.Context, req *v1.UserInfoReq) (*v1.UserInfoRes, error) {
	userId := gconv.String(utility.Auth().GetIdentity(ctx))
	return service.User().GetUserInfo(ctx, userId)
}

/**
我的-修改账户信息
*/
func (c *cUser) UpdateUserInfo(ctx context.Context, req *v1.UpdateUserInfoReq) (*v1.UpdateUserInfoRes, error) {
	userId := gconv.String(utility.Auth().GetIdentity(ctx))
	_, err := service.User().UpdateUserInfo(ctx, userId, req.NickName, req.IntroduceSign)
	if err != nil {
		return nil, err
	}
	return &v1.UpdateUserInfoRes{}, nil
}
