package service

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	v1 "mall-api/api/v1"
	"mall-api/internal/dao"
)

type sAddress struct {
}

var incAddress = sAddress{}

func Address() *sAddress {
	return &incAddress
}

func (s *sAddress) GetUserDefaultAddress(ctx context.Context, userId string) (*v1.AddressRes, error) {
	userAllAddress, err := s.GetUserAllAddress(ctx, userId)
	if err != nil {
		return nil, err
	}
	for _, address := range *userAllAddress {
		if address.DefaultFlag == 1 {
			return &address, err
		}
	}
	return nil, err
}

func (s *sAddress) GetUserAllAddress(ctx context.Context, userId string) (*[]v1.AddressRes, error) {
	var res []v1.AddressRes
	err := dao.UserAddress.Ctx(ctx).Where("user_id=?", userId).Scan(&res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *sAddress) GetAddressById(ctx context.Context, addressId string) (*v1.AddressRes, error) {
	var res v1.AddressRes
	err := dao.UserAddress.Ctx(ctx).Where("address_id=?", addressId).Scan(&res)
	if err != nil {
		return nil, err
	}
	return &res, err
}

func (s *sAddress) AddAddress(ctx context.Context, userId string, req *v1.AddAddressReq) (*v1.AddAddressRes, error) {
	_, err := dao.UserAddress.Ctx(ctx).Insert(g.Map{
		"user_id":        userId,
		"user_name":      req.UserName,
		"user_phone":     req.UserPhone,
		"default_flag":   req.DefaultFlag,
		"province_name":  req.ProvinceName,
		"city_name":      req.CityName,
		"region_name":    req.RegionName,
		"detail_address": req.DetailAddress,
	})
	if err != nil {
		return nil, err
	}
	return &v1.AddAddressRes{}, nil
}
