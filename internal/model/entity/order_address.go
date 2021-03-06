// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// OrderAddress is the golang structure for table order_address.
type OrderAddress struct {
	OrderId       int64  `json:"orderId"       description:""`
	UserName      string `json:"userName"      description:"收货人姓名"`
	UserPhone     string `json:"userPhone"     description:"收货人手机号"`
	ProvinceName  string `json:"provinceName"  description:"省"`
	CityName      string `json:"cityName"      description:"城"`
	RegionName    string `json:"regionName"    description:"区"`
	DetailAddress string `json:"detailAddress" description:"收件详细地址(街道/楼宇/单元)"`
}
