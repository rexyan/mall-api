package controller

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/frame/g"
	v1 "mall-api/api/v1"
	"mall-api/internal/service"
	"mall-api/utility"
)

var Category = cCategory{}

type cCategory struct{}

// GetCategory 获取 GetCategory
func (c *cCategory) GetCategory(ctx context.Context, req *v1.GetCategoryReq) (res *v1.GetCategoryRes, err error) {
	category := service.Category().GetCategory(ctx)
	code := gcode.New(200, "success", category)
	g.RequestFromCtx(ctx).Response.WriteJsonExit(utility.DefaultHandlerRes{
		ResultCode: code.Code(),
		Message:    code.Message(),
		Data:       code.Detail(),
	})
	return
}
