package controller

import (
	"context"
	v1 "mall-api/api/v1"
)

var Goods = cGoods{}

type cGoods struct {}

func (c *cGoods) Detail(ctx context.Context, req v1.GoodsDetailReq) (res *v1.IndexInfosRes, err error)  {
	return nil, err
}