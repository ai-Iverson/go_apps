package controller

import (
	"context"
	"fmt"
	v1 "myapp/api/v1"
	"myapp/internal/service"
)

var (
	User = cUser{}
)

type cUser struct {
}

/*
注册
*/
func (c *cUser) Register(ctx context.Context, req *v1.UserRegisterReq) (res *v1.UserRegisterRes,err error) {
	err1 := service.User().ValidRegist(ctx, req.UserName, req.Email)
	if err1 != nil{
		return nil, err1
	}
	_, err1 = service.User().AddUserToDb(ctx,req.UserName,req.Email,req.PassWord)
	if err1 != nil{
		return nil, err1
	}
	var a = v1.UserRegisterRes{}
	a.Msg = "注册成功"
	return &a, err
}

func (c *cUser) Login(ctx context.Context, req *v1.UserLoginReq) (*v1.UserLoginRes,error) {
	var token *v1.UserLoginRes
	err := service.User().Login(ctx, req.UserName, req.PassWord)
	if err != nil{
		return nil, err
	}
	fmt.Println(token)
	return token, err
}