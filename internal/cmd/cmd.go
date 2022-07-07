package cmd

import (
	"context"
	"mall-api/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"mall-api/internal/controller"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/api/v1", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse, ghttp.MiddlewareCORS)
				group.Bind(
					controller.Index,  // 首页
					controller.User,  // 用户
				)

				group.Group("", func(group *ghttp.RouterGroup) {
					group.Middleware(service.Middleware().Auth)
					group.Bind(
						controller.Goods,   // 商品
					)
				})
			})
			s.Run()
			return nil
		},
	}
)
