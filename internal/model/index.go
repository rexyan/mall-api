package model

type IndexCarouselItemOutput struct {
	CarouselUrl string `json:"carouselUrl"`
	RedirectUrl string `json:"redirectUrl"`
}

type SearchGoodsItem struct {
	GoodsCoverImg string `json:"goodsCoverImg"`
	GoodsId       int    `json:"goodsId"`
	GoodsIntro    int    `json:"goodsIntro"`
	GoodsName     string `json:"goodsName"`
	SellingPrice  int    `json:"sellingPrice"`
}

type IndexSearchOutput struct {
	PageRes
	List []SearchGoodsItem `json:"list"`
}
