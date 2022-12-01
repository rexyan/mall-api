package model

type OrderListItem struct {
	CreateTime             string          `json:"createTime"`
	NewBeeMallOrderItemVOS []GetCartOutput `json:"newBeeMallOrderItemVOS"`
	OrderId                int             `json:"orderId"`
	OrderNo                string          `json:"orderNo"`
	OrderStatus            int             `json:"orderStatus"`
	OrderStatusString      string          `json:"orderStatusString"`
	PayType                int             `json:"payType"`
	TotalPrice             int             `json:"totalPrice"`
}

type OrderListOutput struct {
	PageRes
	List []OrderListItem `json:"list"`
}

type OrderDetailOutput struct {
	OrderListItem
}

type PayOrderOutput struct {
}
