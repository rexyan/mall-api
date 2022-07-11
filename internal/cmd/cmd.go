package cmd

import (
	"context"
	"mall-api/utility"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
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
					controller.Index, // 首页
					controller.User,  // 用户
				)

				group.Group("", func(group *ghttp.RouterGroup) {
					group.Middleware(utility.Middleware().Auth)
					group.Bind(
						controller.Goods,   // 商品
						controller.Cart,    // 购物车
						controller.Address, // 订单地址
						controller.Order,   // 订单
					)
					group.ALLMap(g.Map{
						"/user/info": controller.User.GetUserInfo,
					})
					group.ALLMap(g.Map{
						"/user/info": controller.User.UpdateUserInfo,
					})
				})

			})
			s.Run()
			return nil
		},
	}
)
