package controller

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "myapp/api/v1"
	"myapp/internal/service"
	"myapp/utility"
)

var (
	User = cUser{}
)

type cUser struct {
}

/*
注册
*/
func (c *cUser) Register(ctx context.Context, req *v1.UserRegisterReq) (res *v1.UserRegisterRes, err error) {
	err1 := service.User().ValidRegist(ctx, req.UserName, req.Email)
	if err1 != nil {
		return nil, err1
	}
	_, err1 = service.User().AddUserToDb(ctx, req.UserName, req.Email, req.PassWord)
	if err1 != nil {
		return nil, err1
	}
	var a = v1.UserRegisterRes{}
	a.Msg = "注册成功"
	return &a, err
}

func (c *cUser) Login(ctx context.Context, req *v1.UserLoginReq) (res *v1.UserLoginRes, err error) {
	res = &v1.UserLoginRes{}
	res.Token, res.Expire = utility.Auth().LoginHandler(ctx)
	err = service.User().Login(ctx, req.UserName, req.PassWord)
	if err != nil {
		return nil, err
	}
	return
}

func (c *cUser) RefreshToken(ctx context.Context, req *v1.UserRefreshTokenReq) (res *v1.UserRefreshTokenRes, err error) {
	res = &v1.UserRefreshTokenRes{}
	res.Token, res.Expire = utility.Auth().RefreshHandler(ctx)
	return
}
func (c *cUser) Logout(ctx context.Context, req *v1.UserLogoutReq) (res *v1.UserLogoutRes, err error) {
	utility.Auth().LogoutHandler(ctx)
	return
}

func (c *cUser) UserInfo(ctx context.Context, req *v1.UserInfoReq) (res *v1.UserInfoRes, err error) {
	fmt.Println(gconv.Int(utility.Auth().GetIdentity(ctx)))
	return &v1.UserInfoRes{
		Id:        gconv.Int(utility.Auth().GetIdentity(ctx)),
		IdentiKey: utility.Auth().IdentityKey,
		Payload:   utility.Auth().GetPayload(ctx),
	}, nil
}
