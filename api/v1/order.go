package v1

import "github.com/gogf/gf/v2/frame/g"

/**
下单
*/
type SaveOrderReq struct {
	g.Meta      `path:"saveOrder" tags:"订单" method:"post" summary:"购物车下单"`
	AddressId   int   `json:"addressId" dc:"地址 ID"`
	CartItemIds []int `json:"cartItemIds" dc:"购物车ID"`
}

type SaveOrderRes struct {
	Status string `json:"status"`
}

/**
查询订单状态
*/
type OrderListReq struct {
	PageReq
	Status int `json:"status"`
}

type OrderListRes struct {
	PageRes
	List []struct {
		CreateTime             string   `json:"createTime"`
		NewBeeMallOrderItemVOS []GetCartRes `json:"newBeeMallOrderItemVOS"`
		OrderId                int      `json:"orderId"`
		OrderNo                string   `json:"orderNo"`
		OrderStatus            int      `json:"orderStatus"`
		OrderStatusString      string   `json:"orderStatusString"`
		PayType                int      `json:"payType"`
		TotalPrice             int      `json:"totalPrice"`
	} `json:"list"`
}
