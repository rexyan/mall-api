package index

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"mall-api/internal/dao"
	"mall-api/internal/model"
	"mall-api/internal/model/do"
	"mall-api/internal/service"
)

type (
	sIndex struct{}
)

func init() {
	service.RegisterIndex(New())
}

func New() *sIndex {
	return &sIndex{}
}

// GetCarousels 获取首页轮播图信息
func (s *sIndex) GetCarousels(ctx context.Context) ([]model.IndexCarouselItemOutput, error) {
	var carousels []model.IndexCarouselItemOutput
	err := dao.Carousel.Ctx(ctx).Where(do.Carousel{
		IsDeleted: 0,
	}).Scan(&carousels)
	return carousels, err
}

// GetIndexGoods 获取首页商品信息
func (s *sIndex) GetIndexGoods(ctx context.Context, indexGoodsType int) ([]model.GoodsItemOutput, error) {
	var goods []model.GoodsItemOutput

	err := g.Model("tb_newbee_mall_index_config index_config").
		InnerJoin("tb_newbee_mall_goods_info goods", "index_config.goods_id=goods.goods_id").
		Fields("goods.*").
		Where(g.Map{"index_config.config_type": indexGoodsType}).
		Scan(&goods)
	if err != nil {
		return nil, err
	}
	return goods, err
}

// GetHotGoods 获取首页热门商品信息
func (s *sIndex) GetHotGoods(ctx context.Context) ([]model.GoodsItemOutput, error) {
	return s.GetIndexGoods(ctx, 3)
}

// GetNewGoods 获取首页新商品信息
func (s *sIndex) GetNewGoods(ctx context.Context) ([]model.GoodsItemOutput, error) {
	return s.GetIndexGoods(ctx, 4)
}

// GetRecommendGoods 获取首页推荐商品信息
func (s *sIndex) GetRecommendGoods(ctx context.Context) ([]model.GoodsItemOutput, error) {
	return s.GetIndexGoods(ctx, 5)
}
