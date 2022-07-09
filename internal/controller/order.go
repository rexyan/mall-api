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

func (c cOrder) SaveOrder(ctx context.Context, req *v1.SaveOrderReq) (*v1.SaveOrderRes, error){
	userId := gconv.String(utility.Auth().GetIdentity(ctx))
	err := service.Order().SaveOrder(ctx, userId, req.AddressId, req.CartItemIds)
	if err!=nil{
		return &v1.SaveOrderRes{
			Status: "FAIL",
		}, err
	}
	return &v1.SaveOrderRes{
		Status: "SUCCESS",
	}, err
}