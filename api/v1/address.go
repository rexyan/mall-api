package v1

import "github.com/gogf/gf/v2/frame/g"

// DefaultAddressReq 默认地址
type DefaultAddressReq struct {
	g.Meta `path:"address/default" tags:"地址" method:"get" summary:"默认配送地址"`
}

// UserAllAddressReq 所有地址
type UserAllAddressReq struct {
	g.Meta `path:"address" tags:"地址" method:"get" summary:"所有配送地址"`
}

// AddressDetailReq 地址详情
type AddressDetailReq struct {
	g.Meta    `path:"address/:addressId" tags:"地址" method:"get" summary:"配送地址详情"`
	AddressId string `json:"addressId" in:"path"`
}

type AddressRes struct {
	AddressId     int    `json:"addressId"`
	CityName      string `json:"cityName"`
	CreateTime    string `json:"createTime"`
	DefaultFlag   int    `json:"defaultFlag"`
	DetailAddress string `json:"detailAddress"`
	IsDeleted     int    `json:"isDeleted"`
	ProvinceName  string `json:"provinceName"`
	RegionName    string `json:"regionName"`
	UpdateTime    string `json:"updateTime"`
	UserId        int    `json:"userId"`
	UserName      string `json:"userName"`
	UserPhone     string `json:"userPhone"`
}

// AddAddressReq 新增地址
type AddAddressReq struct {
	g.Meta        `path:"address" tags:"地址" method:"post" summary:"新增配送地址"`
	CityName      string `json:"cityName"`
	DefaultFlag   int    `json:"defaultFlag" d:"0"`
	DetailAddress string `json:"detailAddress"`
	ProvinceName  string `json:"provinceName"`
	RegionName    string `json:"regionName"`
	UserName      string `json:"userName"`
	UserPhone     string `json:"userPhone"`
}

type AddAddressRes struct {
}
