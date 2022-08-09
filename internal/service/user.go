package service

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	v1 "myapp/api/v1"
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

/*
注册
*/
func (s *sUser) ValidRegist(ctx context.Context, username string, mail string) (error) {
	//var msg *v1.UserRegisterRes
	//err :=g.Validator().Data(mail).Messages("邮箱不合法").Run(ctx)
	type BizReq struct {
		MailAddr string `v:"email"`
	}
	var req = BizReq{
		MailAddr: mail,
	}
	err := g.Validator().Data(req).Messages("邮箱输入不合法").Run(ctx)
	if err != nil {
		return err
	}
	r, err1 := g.Model("User").Where(g.Map{
		"username": username,
	}).One()
	if r != nil {
		return gerror.New("用户已存在")
	}
	if err1 != nil {
		return err
	}
	return err
}

func (s *sUser) AddUserToDb(ctx context.Context, UserName, Email,PassWord string) (*v1.UserRegisterRes,error) {
	_,err := g.Model("User").Data(g.Map{
		"username": UserName,
		"email": Email,
		"password": PassWord,
	}).Insert()
	if err != nil{
		return nil, err
	}
	return nil, err
}