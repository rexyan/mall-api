package model

type AddressOutputItem struct {
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

type AddAddressInput struct {
	CityName      string `json:"cityName"`
	DefaultFlag   int    `json:"defaultFlag" d:"0"`
	DetailAddress string `json:"detailAddress"`
	ProvinceName  string `json:"provinceName"`
	RegionName    string `json:"regionName"`
	UserName      string `json:"userName"`
	UserPhone     string `json:"userPhone"`
}
