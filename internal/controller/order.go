package controller

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "mall-api/api/v1"
	"mall-api/internal/model"
	"mall-api/internal/service"
	"mall-api/utility"
)

var Order = cOrder{}

type cOrder struct {
}

// SaveOrder 下单
func (c *cOrder) SaveOrder(ctx context.Context, req *v1.SaveOrderReq) (*v1.SaveOrderRes, error) {
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

// GetOrderByUser 查询用户订单
func (c *cOrder) GetOrderByUser(ctx context.Context, req *v1.OrderListReq) (res *v1.OrderListRes, err error) {
	userId := gconv.String(utility.Auth().GetIdentity(ctx))
	order, err := service.Order().GetOrderByUser(ctx, userId, req.Status, model.PageReq{
		PageNumber: req.PageNumber,
		PageSize:   req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	if err := gconv.Structs(res, &order); err != nil {
		return nil, err
	}
	return
}

// GetOrderDetail 查询订单详情
func (c *cOrder) GetOrderDetail(ctx context.Context, req *v1.OrderDetailReq) (res *v1.OrderDetailRes, err error) {
	orderDetail, err := service.Order().GetOrderDetail(ctx, req.OrderNo)
	if err != nil {
		return nil, err
	}
	if err := gconv.Structs(res, &orderDetail); err != nil {
		return nil, err
	}
	return
}

// PayOrder 订单支付
func (c *cOrder) PayOrder(ctx context.Context, req *v1.PayOrderReq) (res *v1.PayOrderRes, err error) {
	fmt.Println(req)
	order, err := service.Order().PayOrder(ctx, req.OrderNo["data"], req.PayType)
	if err != nil {
		return nil, err
	}
	if err := gconv.Structs(res, &order); err != nil {
		return nil, err
	}
	return
}
