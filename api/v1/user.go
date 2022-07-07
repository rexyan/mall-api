package v1

import "github.com/gogf/gf/v2/frame/g"

type UserLoginReq struct {
	g.Meta      `path:"user/login" tags:"用户" method:"post" summary:"用户登录"`
	LoginName   string `json:"loginName"`
	PasswordMd5 string `json:"passwordMd5"`
}

type UserLoginRes struct {
	Token string `json:"token"`
}
