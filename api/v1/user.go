package v1

import "github.com/gogf/gf/v2/frame/g"

/*
用户注册
*/
type UserRegisterReq struct {
	g.Meta      `path:"user/register" tags:"注册" method:"post" summary:"用户注册"`
	UserName string `json:"username" v:"required|length:4,30#请输入账号|账号长度为:min到:max位"`
	PassWord string `json:"password" v:"required|length:6,30#请输入密码|密码长度不够"`
	Email string `json:"email" v:"email"`
}

type UserRegisterRes struct {
	Msg string
}

type UsersRegisterReq struct {
	g.Meta      `path:"user/register" tags:"注册" method:"post" summary:"用户注册"`
	UserName string `json:"username" v:"required|length:4,30#请输入账号|账号长度为:min到:max位"`
	PassWord string `json:"password" v:"required|length:6,30#请输入密码|密码长度不够"`
	Email string `json:"email" v:"email"`
}

type UsersRegisterRes struct {
	Msg string
}