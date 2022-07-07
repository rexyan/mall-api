package v1

import "github.com/gogf/gf/v2/frame/g"

type GoodsDetailReq struct {
	g.Meta  `path:"goods" tags:"Goods" method:"get" summary:"首页"`
	GoodsId string `json:"goods_id"`
}
