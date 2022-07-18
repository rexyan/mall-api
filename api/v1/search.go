package v1

import "github.com/gogf/gf/v2/frame/g"

/**
商品查询
*/
type IndexSearchReq struct {
	g.Meta          `path:"search" tags:"查询" method:"get" summary:"首页商品查询"`
	Authorization   string `json:"Authorization" in:"header"  dc:"Authorization"`
	Keyword         string `json:"keyword"`
	OrderBy         string `json:"orderBy" d:"new"`
	GoodsCategoryId int    `json:"goodsCategoryId" d:"-1"`
	PageReq
}

type SearchGoodsItem struct {
	GoodsCoverImg string `json:"goodsCoverImg"`
	GoodsId       int    `json:"goodsId"`
	GoodsIntro    int    `json:"goodsIntro"`
	GoodsName     string `json:"goodsName"`
	SellingPrice  int    `json:"sellingPrice"`
}

type IndexSearchRes struct {
	PageRes
	List []SearchGoodsItem `json:"list"`
}
