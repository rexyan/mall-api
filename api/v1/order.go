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
