package controller

import (
	"context"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "mall-api/api/v1"
	"mall-api/internal/service"
	"mall-api/utility"
)

var Cart = cCart{}

type cCart struct{}

/**
获取购物车列表
*/
func (c *cCart) GetShopCart(ctx context.Context, req *v1.GetCartReq) (res *[]v1.GetCartRes, err error) {
	userId := gconv.Int(utility.Auth().GetIdentity(ctx))
	res, err = service.Cart().GetUserCart(ctx, userId)
	if err != nil {
		return nil, err
	}
	return res, err
}

/**
删除购物车中商品
*/
func (c *cCart) DelShopCart(ctx context.Context, req *v1.DelCartReq) (res *v1.DelCartRes, err error) {
	userId := gconv.Int(utility.Auth().GetIdentity(ctx))
	_ = service.Cart().DelShopCart(ctx, userId, req.CarId)
	return nil, nil
}

/**
新增商品到购物车
*/
func (c *cCart) AddShopCart(ctx context.Context, req *v1.AddCartReq) (res *v1.AddCartRes, err error) {
	userId := gconv.Int(utility.Auth().GetIdentity(ctx))
	res, err = service.Cart().AddShopCart(ctx, userId, req.GoodsId, req.GoodsCount)
	if err != nil {
		return nil, err
	}
	return res, err
}

/**
修改购物车中商品数量
*/
func (c *cCart) UpdateShopCart(ctx context.Context, req *v1.UpdateCartReq) (res *v1.UpdateCartRes, err error) {
	userId := gconv.Int(utility.Auth().GetIdentity(ctx))
	res, err = service.Cart().UpdateShopCart(ctx, userId, req.CartItemId, req.GoodsCount)
	if err != nil {
		return nil, err
	}
	return res, err
}

/**
购物车结算
*/
func (c *cCart) CartSettle(ctx context.Context, req *v1.CartSettleReq) (res []v1.CartSettleRes, err error) {
	cartItemIds := gstr.Split(req.CartItemIds, ",")
	return service.Cart().CartSettle(ctx, cartItemIds)
}
