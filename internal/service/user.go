// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package service

import (
	"context"
	v1 "mall-api/api/v1"
	"mall-api/internal/model/entity"
)

type IUser interface {
	Login(ctx context.Context, username string, password string) (err error)
	CheckUserPassword(ctx context.Context, username string, password string) (user *entity.User)
	GetUserById(ctx context.Context, userId string) *entity.User
	GetUserInfo(ctx context.Context, userId string) (*v1.UserInfoRes, error)
	UpdateUserInfo(ctx context.Context, userId, nickName, introduceSign string) (bool, error)
}

var localUser IUser

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
