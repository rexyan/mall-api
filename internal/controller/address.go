package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "mall-api/api/v1"
	"mall-api/internal/model"
	"mall-api/internal/service"
	"mall-api/utility"
)

var Address = cAddress{}

type cAddress struct{}

// DefaultAddress 默认地址
func (c *cAddress) DefaultAddress(ctx context.Context, req *v1.DefaultAddressReq) (res *v1.AddressRes, err error) {
	userId := gconv.String(utility.Auth().GetIdentity(ctx))
	userDefaultAddress, err := service.Address().GetUserDefaultAddress(ctx, userId)
	if err != nil {
		return nil, err
	}
	if err := gconv.Structs(userDefaultAddress, &res); err != nil {
		return nil, err
	}
	return
}

// AddressDetail 地址详情
func (c *cAddress) AddressDetail(ctx context.Context, req *v1.AddressDetailReq) (res *v1.AddressRes, err error) {
	address, err := service.Address().GetAddressById(ctx, req.AddressId)
	if err != nil {
		return nil, err
	}
	if err := gconv.Structs(address, &res); err != nil {
		return nil, err
	}
	return
}

// GetUserAllAddress 获取用户所有地址
func (c *cAddress) GetUserAllAddress(ctx context.Context, req *v1.UserAllAddressReq) (res []*v1.AddressRes, err error) {
	userId := gconv.String(utility.Auth().GetIdentity(ctx))
	address, err := service.Address().GetUserAllAddress(ctx, userId)
	if err != nil {
		return nil, err
	}
	if err := gconv.Structs(address, &res); err != nil {
		return nil, err
	}
	return
}

// AddAddress 新增地址
func (c *cAddress) AddAddress(ctx context.Context, req *v1.AddAddressReq) (res *v1.AddAddressRes, err error) {
	userId := gconv.String(utility.Auth().GetIdentity(ctx))
	address, err := service.Address().AddAddress(ctx, userId, &model.AddAddressInput{
		CityName:      req.CityName,
		DefaultFlag:   req.DefaultFlag,
		DetailAddress: req.DetailAddress,
		ProvinceName:  req.ProvinceName,
		RegionName:    req.RegionName,
		UserName:      req.UserName,
		UserPhone:     req.UserPhone,
	})
	if err != nil {
		return nil, err
	}
	if err := gconv.Structs(address, &res); err != nil {
		return nil, err
	}
	return
}
