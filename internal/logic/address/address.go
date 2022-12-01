package address

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"mall-api/internal/dao"
	"mall-api/internal/model"
	"mall-api/internal/service"
)

type sAddress struct {
}

func init() {
	service.RegisterAddress(New())
}

func New() *sAddress {
	return &sAddress{}
}

func (s *sAddress) GetUserDefaultAddress(ctx context.Context, userId string) (*model.AddressOutputItem, error) {
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

func (s *sAddress) GetUserAllAddress(ctx context.Context, userId string) (*[]model.AddressOutputItem, error) {
	var res []model.AddressOutputItem
	err := dao.UserAddress.Ctx(ctx).Where("user_id=?", userId).Scan(&res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *sAddress) GetAddressById(ctx context.Context, addressId string) (*model.AddressOutputItem, error) {
	var res model.AddressOutputItem
	err := dao.UserAddress.Ctx(ctx).Where("address_id=?", addressId).Scan(&res)
	if err != nil {
		return nil, err
	}
	return &res, err
}

func (s *sAddress) AddAddress(ctx context.Context, userId string, req *model.AddAddressInput) (*model.AddressOutputItem, error) {
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
	return &model.AddressOutputItem{}, nil
}
