package controller

import (
	"context"
	v1 "mall-api/api/v1"
	"mall-api/internal/service"
)

var Category = cCategory{}

type cCategory struct{}

func (c *cCategory) GetCategory(ctx context.Context, req *v1.CategoryReq) (res *[]v1.CategoryRes, err error) {
	return service.Category().GetCategory(ctx)
}
