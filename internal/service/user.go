package service

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	v1 "myapp/api/v1"
	"myapp/internal/model/entity"
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
	//encryptPassword, _ := bcrypt.GenerateFromPassword([]byte(PassWord), 10)
	_, err := g.Model("User").Data(g.Map{
		"username":   UserName,
		"email":      Email,
		"password":   PassWord,
		"is_deleted": 0,
		"active":     1,
	}).Insert()
	if err != nil {
		return nil, err
	}
	return nil, err
}
func (s *sUser) CheckUserPassword(ctx context.Context, username string, password string) (user *entity.User) {
	var u  *entity.User
	err := g.Model("user").Where(g.Map{
		"username":   username,
		"password":   password,
		"is_deleted": 0,
	}).Scan(u)
	if err != nil {
		return nil
	}
	//err1 := bcrypt.CompareHashAndPassword([]byte(record["password"].String()), []byte(password))
	return u
}

func (s *sUser) Login(ctx context.Context, username string, password string) (err error) {
	record,err := g.Model("user").Where(g.Map{
		"username":   username,
		"password":   password,
		"is_deleted": 0,
	}).One()
	fmt.Print(record)
	if err != nil && g.IsNil(record){
		return err
	}
	if err != nil{
		return err
	}
	return nil
}