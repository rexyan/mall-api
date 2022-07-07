package controller

import (
	"context"
	v1 "mall-api/api/v1"
	"mall-api/internal/service"
)

var Goods = cGoods{}

type cGoods struct {}

func (c *cGoods) Detail(ctx context.Context, req *v1.GoodsDetailReq) (res *v1.GoodsDetailRes, err error)  {
	res, err = service.Goods().Detail(ctx, req.GoodsId)
	if err != nil {
		return nil, err
	}
	return res, err
}

