package v1

import "github.com/gogf/gf/v2/frame/g"

type GoodsDetailReq struct {
	g.Meta        `path:"goods/detail/:goods_id" tags:"商品" method:"get" summary:"商品详情"`
	GoodsId       string `json:"goods_id" in:"path"  dc:"商品ID"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
}

type GoodsDetailRes struct {
	GoodsCarouselList  []string `json:"goodsCarouselList"`
	GoodsCoverImg      string   `json:"goodsCoverImg"`
	GoodsDetailContent string   `json:"goodsDetailContent"`
	GoodsId            int      `json:"goodsId"`
	GoodsIntro         string   `json:"goodsIntro"`
	GoodsName          string   `json:"goodsName"`
	OriginalPrice      int      `json:"originalPrice"`
	SellingPrice       int      `json:"sellingPrice"`
	Tag                string   `json:"tag"`
}
