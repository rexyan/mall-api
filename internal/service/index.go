package service

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	v1 "mall-api/api/v1"
	"mall-api/internal/dao"
	"mall-api/internal/model/do"
)

type (
	sIndex struct{}
)

var (
	insIndex = sIndex{}
)

func Index() *sIndex {
	return &insIndex
}

/**
获取首页轮播图信息
*/
func (s *sIndex) GetCarousels(ctx context.Context) ([]v1.IndexCarouselItem, error) {
	var carousels []v1.IndexCarouselItem
	err := dao.Carousel.Ctx(ctx).Where(do.Carousel{
		IsDeleted: 0,
	}).Scan(&carousels)
	return carousels, err
}

/**
获取首页商品信息
*/
func (s *sIndex) GetIndexGoods(ctx context.Context, indexGoodsType int) ([]v1.GoodsItem, error) {
	var goods []v1.GoodsItem

	err := g.Model("tb_newbee_mall_index_config index_config").InnerJoin("tb_newbee_mall_goods_info goods", "index_config.goods_id=goods.goods_id").Fields(
		"goods.*",
		).Where(g.Map{"index_config.config_type": indexGoodsType}).Scan(&goods)
	if err != nil {
		return nil, err
	}
	return goods, err
}

/**
获取首页热门商品信息
*/
func (s *sIndex) GetHotGoods(ctx context.Context) ([]v1.GoodsItem, error) {
	return s.GetIndexGoods(ctx, 3)
}

/**
获取首页新商品信息
*/
func (s *sIndex) GetNewGoods(ctx context.Context) ([]v1.GoodsItem, error) {
	return s.GetIndexGoods(ctx, 4)
}

/**
获取首页推荐商品信息
*/
func (s *sIndex) GetRecommendGoods(ctx context.Context) ([]v1.GoodsItem, error) {
	return s.GetIndexGoods(ctx, 5)
}
