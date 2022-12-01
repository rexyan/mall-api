package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "mall-api/api/v1"
	"mall-api/internal/service"
)

var Goods = cGoods{}

type cGoods struct{}

func (c *cGoods) Detail(ctx context.Context, req *v1.GoodsDetailReq) (res *v1.GoodsDetailRes, err error) {
	goodsDetail, err := service.Goods().Detail(ctx, req.GoodsId)
	if err != nil {
		return nil, err
	}
	if err := gconv.Structs(goodsDetail, &res); err != nil {
		return nil, err
	}
	return
}
