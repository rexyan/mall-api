package v1

import "github.com/gogf/gf/v2/frame/g"

/**
默认地址
*/
type DefaultAddressReq struct {
	g.Meta `path:"address/default" tag:"地址" method:"get" summary:"默认配送地址"`
}

/**
所有地址
*/
type UserAllAddressReq struct {
	g.Meta `path:"address" tag:"地址" method:"get" summary:"所有配送地址"`
}

/**
地址详情
*/
type AddressDetailReq struct {
	g.Meta `path:"address/:addressId" tag:"地址" method:"get" summary:"配送地址详情"`
	AddressId string `json:"addressId" in:"path"`
}

type AddressRes struct {
	AddressId     int    `json:"address_id"`
	CityName      string `json:"city_name"`
	CreateTime    string `json:"create_time"`
	DefaultFlag   int    `json:"default_flag"`
	DetailAddress string `json:"detail_address"`
	IsDeleted     int    `json:"is_deleted"`
	ProvinceName  string `json:"province_name"`
	RegionName    string `json:"region_name"`
	UpdateTime    string `json:"update_time"`
	UserId        int    `json:"user_id"`
	UserName      string `json:"user_name"`
	UserPhone     string `json:"user_phone"`
}
