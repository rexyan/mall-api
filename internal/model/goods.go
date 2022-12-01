package model

type GoodsDetailOutput struct {
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

type GoodsItemOutput struct {
	GoodsCoverImg string `json:"goodsCoverImg"`
	GoodsId       uint   `json:"goodsId"`
	GoodsIntro    string `json:"goodsIntro"`
	GoodsName     string `json:"goodsName"`
	SellingPrice  uint   `json:"sellingPrice"`
	Tag           string `json:"tag"`
}
