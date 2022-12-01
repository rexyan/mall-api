package v1

import "github.com/gogf/gf/v2/frame/g"

/*
*
下单
*/
type SaveOrderReq struct {
	g.Meta      `path:"saveOrder" tags:"订单" method:"post" summary:"购物车下单"`
	AddressId   int   `json:"addressId" dc:"地址 ID"`
	CartItemIds []int `json:"cartItemIds" dc:"购物车ID"`
}

type SaveOrderRes struct {
	Data string `json:"data"`
}

/*
*
查询订单状态
*/
type OrderListReq struct {
	g.Meta        `path:"order" tags:"订单" method:"get" summary:"用户订单查询"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	Status        int    `json:"status" d:"-1"`
	PageReq
}

type OrderListItem struct {
	CreateTime             string       `json:"createTime"`
	NewBeeMallOrderItemVOS []GetCartRes `json:"newBeeMallOrderItemVOS"`
	OrderId                int          `json:"orderId"`
	OrderNo                string       `json:"orderNo"`
	OrderStatus            int          `json:"orderStatus"`
	OrderStatusString      string       `json:"orderStatusString"`
	PayType                int          `json:"payType"`
	TotalPrice             int          `json:"totalPrice"`
}

type OrderListRes struct {
	PageRes
	List []OrderListItem `json:"list"`
}

// OrderDetailReq 查询订单详情
type OrderDetailReq struct {
	g.Meta        `path:"order/:orderNo" tags:"订单" method:"get" summary:"查询订单详情"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	OrderNo       string `json:"orderNo" in:"path"`
}

type OrderDetailRes struct {
	OrderListItem
}

// PayOrderReq 订单支付
type PayOrderReq struct {
	g.Meta        `path:"paySuccess" tags:"订单" method:"get" summary:"订单支付"`
	Authorization string            `json:"Authorization" in:"header"  dc:"Authorization"`
	OrderNo       map[string]string `json:"orderNo" in:"query"`
	PayType       string            `json:"payType" in:"query"`
}

type PayOrderRes struct {
}
