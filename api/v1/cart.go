package v1

import "github.com/gogf/gf/v2/frame/g"

/**
获取购物车列表
*/
type GetCartReq struct {
	g.Meta        `path:"shop-cart" tag:"商品" method:"get" summary:"获取购物车商品"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
}

type GetCartRes struct {
	CartItemId    int    `json:"cartItemId"`
	GoodsCount    int    `json:"goodsCount"`
	GoodsCoverImg string `json:"goodsCoverImg"`
	GoodsName     string `json:"goodsName"`
	SellingPrice  int    `json:"sellingPrice"`
}

/**
删除购物车中商品
*/
type DelCartReq struct {
	g.Meta        `path:"shop-cart/:cart_id" tag:"商品" method:"delete" summary:"删除购物车商品"`
	CarId         string `json:"cart_id" in:"path"  dc:"购物车ID"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
}

type DelCartRes struct {
}

/**
新增商品到购物车
*/
type AddCartReq struct {
	g.Meta        `path:"shop-cart" tag:"商品" method:"post" summary:"添加购物车商品"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	GoodsCount    int    `json:"goodsCount" dc:"数量"`
	GoodsId       int    `json:"goodsId" dc:"商品ID"`
}

type AddCartRes struct {
}

/**
修改购物车中商品数量
*/
type UpdateCartReq struct {
	g.Meta        `path:"shop-cart" tag:"商品" method:"put" summary:"修改购物车中商品数量"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	CartItemId    int    `json:"cartItemId" dc:"购物车ID"`
	GoodsCount    int    `json:"goodsId" dc:"商品数量"`
}

type UpdateCartRes struct {
}
