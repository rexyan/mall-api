package user

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "mall-api/api/v1"
	"mall-api/internal/dao"
	"mall-api/internal/model/entity"
	"mall-api/internal/service"
)

type (
	sUser struct{}
)

func init() {
	service.RegisterUser(New())
}

func New() *sUser {
	return &sUser{}
}

func (s *sUser) Login(ctx context.Context, username string, password string) (err error) {
	userCls := dao.User.Columns()
	record, err := dao.User.Ctx(ctx).
		Where(userCls.LoginName, username).
		Where(userCls.PasswordMd5, password).
		Where(userCls.IsDeleted, 0).
		One()
	if err != nil && g.IsNil(record) {
		return err
	}
	return nil
}

func (s *sUser) CheckUserPassword(ctx context.Context, username string, password string) (user *entity.User) {
	userCls := dao.User.Columns()
	err := dao.User.Ctx(ctx).
		Where(userCls.LoginName, username).
		Where(userCls.PasswordMd5, password).
		Where(userCls.IsDeleted, 0).
		Scan(&user)
	if err != nil {
		return nil
	}
	return
}

// GetUserById 获取用户
func (s *sUser) GetUserById(ctx context.Context, userId string) *entity.User {
	var user *entity.User
	err := dao.User.Ctx(ctx).Where(dao.User.Columns().UserId, userId).Scan(&user)
	if err != nil {
		return nil
	}
	return user
}

// GetUserInfo 获取用户信息
func (s *sUser) GetUserInfo(ctx context.Context, userId string) (*v1.UserInfoRes, error) {
	var userInfo *v1.UserInfoRes
	user := s.GetUserById(ctx, userId)
	if user == nil {
		return nil, gerror.New("user not fund!")
	}
	err := gconv.Scan(user, &userInfo)
	if err != nil {
		return nil, gerror.New("user not fund!")
	}
	return userInfo, nil
}

// UpdateUserInfo 我的-修改账户信息
func (s *sUser) UpdateUserInfo(ctx context.Context, userId, nickName, introduceSign string) (bool, error) {
	userCls := dao.User.Columns()
	_, err := dao.User.Ctx(ctx).Data(g.Map{
		userCls.NickName:      nickName,
		userCls.IntroduceSign: introduceSign,
	}).Where(userCls.UserId, userId).Update()
	if err != nil {
		return false, err
	}
	return true, nil
}
