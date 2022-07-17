package controller

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"mall-api/internal/service"
)

func GetCategory(r *ghttp.Request) {
	ctx := context.TODO()
	res := service.Category().GetCategory(ctx)
	r.Response.WriteJson(g.Map{
		"data":       res,
		"message":    "",
		"resultCode": 200,
	})
}
