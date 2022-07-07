package service

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
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


func (s *sUser) GetUserInfo(ctx context.Context, username string, password string) (user *entity.User)  {
	var u *entity.User
	err := g.Model("tb_newbee_mall_user").Where(g.Map{
		"login_name":   username,
		"password_md5": password,
		"is_deleted":   0,
	}).Scan(&u)

	if err != nil{
		g.Log("xxx")
		return nil
	}
	return u
}
