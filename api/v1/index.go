package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type IndexCarouselItem struct {
	CarouselUrl string `json:"carouselUrl"`
	RedirectUrl string `json:"redirectUrl"`
}

type GoodsItem struct {
	GoodsCoverImg string `json:"goodsCoverImg"`
	GoodsId       uint   `json:"goodsId"`
	GoodsIntro    string `json:"goodsIntro"`
	GoodsName     string `json:"goodsName"`
	SellingPrice  uint   `json:"sellingPrice"`
	Tag           string `json:"tag"`
}

// 首页
type IndexInfosReq struct {
	g.Meta `path:"/api/v1/index-infos" tags:"Index" method:"get" summary:"首页"`
}

type IndexInfosRes struct {
	IndexCarouselItem   []IndexCarouselItem `json:"carousels"`
	IndexHotGoods       []GoodsItem         `json:"hotGoodses"`
	IndexNewGoods       []GoodsItem         `json:"newGoodses"`
	IndexRecommendGoods []GoodsItem         `json:"recommendGoodses"`
}
