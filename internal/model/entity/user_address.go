// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserAddress is the golang structure for table user_address.
type UserAddress struct {
	AddressId     int64       `json:"addressId"     description:""`
	UserId        int64       `json:"userId"        description:"用户主键id"`
	UserName      string      `json:"userName"      description:"收货人姓名"`
	UserPhone     string      `json:"userPhone"     description:"收货人手机号"`
	DefaultFlag   int         `json:"defaultFlag"   description:"是否为默认 0-非默认 1-是默认"`
	ProvinceName  string      `json:"provinceName"  description:"省"`
	CityName      string      `json:"cityName"      description:"城"`
	RegionName    string      `json:"regionName"    description:"区"`
	DetailAddress string      `json:"detailAddress" description:"收件详细地址(街道/楼宇/单元)"`
	IsDeleted     int         `json:"isDeleted"     description:"删除标识字段(0-未删除 1-已删除)"`
	CreateTime    *gtime.Time `json:"createTime"    description:"添加时间"`
	UpdateTime    *gtime.Time `json:"updateTime"    description:"修改时间"`
}
