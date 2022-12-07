package model

type UserCartOutput struct {
	CartItemId    int    `json:"cartItemId"`
	GoodsId       int    `json:"goodsId"`
	GoodsCount    int    `json:"goodsCount"`
	GoodsCoverImg string `json:"goodsCoverImg"`
	GoodsName     string `json:"goodsName"`
	SellingPrice  int    `json:"sellingPrice"`
}

type AddCartOutput struct {
}

type UpdateCartOutput struct {
}

type CartSettleOutput struct {
	UserCartOutput
	GoodsId int `json:"goodsId"`
}

type GetCartOutput struct {
	CartItemId    int    `json:"cartItemId"`
	GoodsId       int    `json:"goodsId"`
	GoodsCount    int    `json:"goodsCount"`
	GoodsCoverImg string `json:"goodsCoverImg"`
	GoodsName     string `json:"goodsName"`
	SellingPrice  int    `json:"sellingPrice"`
}

type GetUserCartFields struct {
	CartItemId    int    `json:"cartItemId"`
	GoodsCount    int    `json:"goodsCount"`
	GoodsCoverImg string `json:"goodsCoverImg"`
	GoodsName     string `json:"goodsName"`
	SellingPrice  int    `json:"sellingPrice"`
}
