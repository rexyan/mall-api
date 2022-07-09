package service

import (
	"context"
	v1 "mall-api/api/v1"
	"mall-api/internal/dao"
)

type sAddress struct {

}

var incAddress = sAddress{}

func Address() * sAddress{
	return &incAddress
}

func (s *sAddress) GetUserDefaultAddress(ctx context.Context, userId string) (*v1.AddressRes, error) {
	userAllAddress, err := s.GetUserAllAddress(ctx, userId)
	if err!=nil{
		return nil, err
	}
	for _, address:=range *userAllAddress{
		if address.DefaultFlag == 1{
			return &address, err
		}
	}
	return nil, err
}

func (s *sAddress) GetUserAllAddress(ctx context.Context, userId string) (*[]v1.AddressRes, error) {
	var res []v1.AddressRes
	err := dao.UserAddress.Ctx(ctx).Where("user_id=?", userId).Scan(&res)
	if err!=nil{
		return nil, err
	}
	return &res, nil
}

func (s *sAddress) GetAddressById(ctx context.Context, addressId string) (*v1.AddressRes, error) {
	var res v1.AddressRes
	err := dao.UserAddress.Ctx(ctx).Where("address_id=?", addressId).Scan(&res)
	if err!=nil{
		return nil, err
	}
	return &res, err
}