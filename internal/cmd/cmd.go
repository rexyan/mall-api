package cmd

import (
	"context"
	"mall-api/utility"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"mall-api/internal/controller"
	_ "mall-api/internal/logic"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			//全局中间件注册
			s.Use(utility.Middleware().CustomResponse, ghttp.MiddlewareCORS)

			// 不需要 JWT Token 校验的路由
			s.Group("/api/v1", func(group *ghttp.RouterGroup) {
				group.GET("/index-infos", controller.Index.IndexInfos)
				group.POST("/user/login", controller.User.Login)
				group.POST("/categories", controller.Category.GetCategory)

				// 需要  JWT Token 的路由
				group.Group("", func(group *ghttp.RouterGroup) {
					group.Middleware(utility.Middleware().Auth)
					group.Bind(
						controller.Goods,               // 商品
						controller.Cart,                // 购物车
						controller.Address,             // 订单地址
						controller.Order,               // 订单
						controller.Search,              // 查询
						controller.User.GetUserInfo,    // 用户
						controller.User.UpdateUserInfo, // 用户
					)
				})
			})
			s.Run()
			return nil
		},
	}
)
