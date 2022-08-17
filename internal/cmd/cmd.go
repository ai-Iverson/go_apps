package cmd

import (
	"context"
	"myapp/utility"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"myapp/internal/controller"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Use(utility.Middleware().CustomResponse, ghttp.MiddlewareCORS)
			s.Group("/api/v1", func(group *ghttp.RouterGroup) {
				group.POST("/user/register", controller.User.Register)
				group.POST("/user/login", controller.User.Login)
				group.POST("/user/user/refresh_token", controller.User.RefreshToken)
				group.POST("/user/user/logout", controller.User.Logout)
			})
			s.Group("/api/v1", func(group *ghttp.RouterGroup) {
				group.Middleware(utility.Middleware().Auth)
				group.GET("/user/getuserinfo", controller.User.UserInfo)

			})
			_ = controller.Schedules.Initialize()
			s.Run()
			return nil
		},
	}
)
