package v1

import "github.com/gogf/gf/v2/frame/g"

// UserLoginReq 用户登录
type UserLoginReq struct {
	g.Meta      `path:"user/login" tags:"用户" method:"post" summary:"用户登录"`
	LoginName   string `json:"loginName"`
	PasswordMd5 string `json:"passwordMd5"`
}

type UserLoginRes string

// UserInfoReq 我的-用户信息
type UserInfoReq struct {
	g.Meta        `path:"user/info" tags:"用户" method:"get" summary:"用户信息"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
}

type UserInfoRes struct {
	IntroduceSign string `json:"introduceSign"`
	LoginName     string `json:"loginName"`
	NickName      string `json:"nickName"`
}

// UpdateUserInfoReq 我的-修改账户信息
type UpdateUserInfoReq struct {
	g.Meta        `path:"user/info" tags:"用户" method:"put" summary:"修改用户信息"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	IntroduceSign string `json:"introduceSign"`
	NickName      string `json:"nickName"`
}

type UpdateUserInfoRes struct{}
