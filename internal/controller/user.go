package controller

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
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
func Login(r *ghttp.Request) {
	ctx := r.GetCtx()
	name := gconv.String(r.Get("loginName"))
	pass := gconv.String(r.Get("passwordMd5"))

	if err := service.User().Login(ctx, name, pass); err != nil {
		fmt.Println(err)
		r.Response.WriteJson(g.Map{
			"data":       "",
			"message":    "用户密码不正确",
			"resultCode": 401,
		})
	}
	tokenString, _ := utility.Auth().LoginHandler(ctx)
	r.Response.WriteJson(g.Map{
		"data":       tokenString,
		"message":    "success",
		"resultCode": 200,
	})
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
