package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "mall-api/api/v1"
	"mall-api/internal/service"
	"mall-api/utility"
)

var Order = cOrder{}

type cOrder struct {
}

/**
下单
*/
func (c cOrder) SaveOrder(ctx context.Context, req *v1.SaveOrderReq) (*v1.SaveOrderRes, error) {
	userId := gconv.String(utility.Auth().GetIdentity(ctx))
	orderNo, err := service.Order().SaveOrder(ctx, userId, req.AddressId, req.CartItemIds)
	if err != nil {
		return &v1.SaveOrderRes{
			Data: "",
		}, err
	}
	return &v1.SaveOrderRes{
		Data: orderNo,
	}, err
}

/**
查询用户订单
*/
func (c cOrder) GetOrderByUser(ctx context.Context, req *v1.OrderListReq) (*v1.OrderListRes, error) {
	userId := gconv.String(utility.Auth().GetIdentity(ctx))
	return service.Order().GetOrderByUser(ctx, userId, req.Status, req.PageReq)
}

/**
查询订单详情
*/
func (c *cOrder) GetOrderDetail(ctx context.Context, req *v1.OrderDetailReq) (*v1.OrderDetailRes, error) {
	return service.Order().GetOrderDetail(ctx, req.OrderNo)
}

/**
订单支付
 */
func (c *cOrder) PayOrder(ctx context.Context, req *v1.PayOrderReq) (*v1.PayOrderRes, error) {
	return service.Order().PayOrder(ctx, req.OrderNo, req.PayType)
}