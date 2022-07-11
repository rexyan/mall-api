package service

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "mall-api/api/v1"
	"mall-api/internal/dao"
	"mall-api/internal/model/entity"
)

type (
	sUser struct{}
)

var (
	insUser = sUser{}
)

func User() *sUser {
	return &insUser
}

/**
Login
*/
func (s *sUser) Login(ctx context.Context, username string, password string) (err error) {
	record, err := g.Model("tb_newbee_mall_user").Where(g.Map{
		"login_name":   username,
		"password_md5": password,
		"is_deleted":   0,
	}).One()

	if err != nil && g.IsNil(record) {
		return err
	}
	return nil
}

func (s *sUser) CheckUserPassword(ctx context.Context, username string, password string) (user *entity.User) {
	var u *entity.User
	err := g.Model("tb_newbee_mall_user").Where(g.Map{
		"login_name":   username,
		"password_md5": password,
		"is_deleted":   0,
	}).Scan(&u)

	if err != nil {
		g.Log("xxx")
		return nil
	}
	return u
}

/**
获取用户
*/
func (s *sUser) GetUserById(ctx context.Context, userId string) *entity.User {
	var user *entity.User
	err := dao.User.Ctx(ctx).Where("user_id", userId).Scan(&user)
	if err != nil {
		return nil
	}
	return user
}

/**
获取用户信息
*/
func (s *sUser) GetUserInfo(ctx context.Context, userId string) (*v1.UserInfoRes, error) {
	var userInfo *v1.UserInfoRes
	user := User().GetUserById(ctx, userId)
	if user == nil {
		return nil, gerror.New("user not fund!")
	}
	err := gconv.Scan(user, &userInfo)
	if err != nil {
		return nil, gerror.New("user not fund!")
	}
	return userInfo, nil
}

/**
我的-修改账户信息
*/
func (s *sUser) UpdateUserInfo(ctx context.Context, userId, nickName, introduceSign string) (bool, error) {
	_, err := dao.User.Ctx(ctx).Data(g.Map{
		"nick_name": nickName,
		"introduce_sign": introduceSign,
	}).Where("user_id", userId).Update()
	if err !=nil{
		return false, err
	}
	return true, nil
}
