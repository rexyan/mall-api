package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "mall-api/api/v1"
	"mall-api/internal/service"
	"mall-api/utility"
)

var Address = cAddress{}

type cAddress struct{}

/**
默认地址
*/
func (c *cAddress) DefaultAddress(ctx context.Context, req *v1.DefaultAddressReq) (*v1.AddressRes, error) {
	userId := gconv.String(utility.Auth().GetIdentity(ctx))
	return service.Address().GetUserDefaultAddress(ctx, userId)
}

/**
地址详情
*/
func (c *cAddress) AddressDetail(ctx context.Context, req *v1.AddressDetailReq) (*v1.AddressRes, error) {
	return service.Address().GetAddressById(ctx, req.AddressId)
}

/**
地址详情
*/
func (c *cAddress) GetUserAllAddress(ctx context.Context, req *v1.UserAllAddressReq) (*[]v1.AddressRes, error) {
	userId := gconv.String(utility.Auth().GetIdentity(ctx))
	return service.Address().GetUserAllAddress(ctx, userId)
}

/**
新增地址
 */
func (c *cAddress) AddAddress (ctx context.Context, req *v1.AddAddressReq) (*v1.AddAddressRes, error) {
	userId := gconv.String(utility.Auth().GetIdentity(ctx))
	return service.Address().AddAddress(ctx, userId, req)
}
