package service

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/crypto/gmd5"
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
func (s *sUser) ValidRegist(ctx context.Context, username string, mail string) error {
	//var msg *v1.UserRegisterRes
	//err :=g.Validator().Data(mail).Messages("邮箱不合法").Run(ctx)

	r, err1 := g.Model("User").Where(g.Map{
		"username": username,
	}).One()
	if r != nil {
		return gerror.New("用户已存在")
	}
	if err1 != nil {
		return err1
	}
	return err1
}

func (s *sUser) AddUserToDb(ctx context.Context, UserName, Email, PassWord string) (*v1.UserRegisterRes, error) {
	encryptPassword, errpass := gmd5.Encrypt(PassWord + "qwer")
	if errpass != nil {
		return nil, errors.New("服务器异常")
	}
	_, err := g.Model("User").Data(g.Map{
		"username":   UserName,
		"email":      Email,
		"password":   encryptPassword,
		"is_deleted": 0,
		"active":     1,
	}).Insert()
	if err != nil {
		return nil, err
	}
	return nil, err
}
func (s *sUser) CheckUserPassword(ctx context.Context, username string, password string) map[string]interface{} {
	encryptPassword, errpass := gmd5.Encrypt(password + "qwer")
	if errpass != nil {
		return nil
	}
	user,usererr := g.Model("user").Where(g.Map{
		"username":   username,
		"password":   encryptPassword,
		"is_deleted": 0,
	}).One()
	if usererr != nil {
		return nil
	}
	return g.Map{
		"id":user["id"],
		"username":user["username"],
	}
}

func (s *sUser) Login(ctx context.Context, username string, password string) (err error) {
	encryptPassword, errpass := gmd5.Encrypt(password + "qwer")
	if errpass != nil {
		return errors.New("服务器异常")
	}
	record, err := g.Model("user").Where(g.Map{
		"username":   username,
		"password":   encryptPassword,
		"is_deleted": 0,
	}).One()
	if err != nil && g.IsNil(record) {
		return err
	}
	if err != nil {
		return err
	}
	return nil
}
