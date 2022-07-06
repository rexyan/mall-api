// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Order is the golang structure for table order.
type Order struct {
	OrderId     int64       `json:"orderId"     description:"订单表主键id"`
	OrderNo     string      `json:"orderNo"     description:"订单号"`
	UserId      int64       `json:"userId"      description:"用户主键id"`
	TotalPrice  int         `json:"totalPrice"  description:"订单总价"`
	PayStatus   int         `json:"payStatus"   description:"支付状态:0.未支付,1.支付成功,-1:支付失败"`
	PayType     int         `json:"payType"     description:"0.无 1.支付宝支付 2.微信支付"`
	PayTime     *gtime.Time `json:"payTime"     description:"支付时间"`
	OrderStatus int         `json:"orderStatus" description:"订单状态:0.待支付 1.已支付 2.配货完成 3:出库成功 4.交易成功 -1.手动关闭 -2.超时关闭 -3.商家关闭"`
	ExtraInfo   string      `json:"extraInfo"   description:"订单body"`
	IsDeleted   int         `json:"isDeleted"   description:"删除标识字段(0-未删除 1-已删除)"`
	CreateTime  *gtime.Time `json:"createTime"  description:"创建时间"`
	UpdateTime  *gtime.Time `json:"updateTime"  description:"最新修改时间"`
}
