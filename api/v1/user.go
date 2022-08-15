package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

/*
用户注册
*/
type UserRegisterReq struct {
	g.Meta   `path:"user/register" tags:"注册" method:"post" summary:"用户注册"`
	UserName string `json:"username" v:"required|length:4,30#请输入账号|账号长度为:4到:30位"`
	PassWord string `json:"password" v:"required|length:6,30#请输入密码|密码长度不够"`
	Email    string `json:"email" v:"email"`
}

type UserRegisterRes struct {
	Msg string
}

type UserLoginReq struct {
	g.Meta   `path:"user/login" tags:"登录" method:"post" summary:"用户登录"`
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

type UserLoginRes struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}

type UserRefreshTokenReq struct {
	g.Meta `path:"user/refresh_token" tags:"刷新Token" method:"post" summary:"刷新token"`
}

type UserRefreshTokenRes struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}

type UserLogoutReq struct {
	g.Meta `path:"user/logout" tags:"登出" method:"post" summary:"登出"`
}

type UserLogoutRes struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}
