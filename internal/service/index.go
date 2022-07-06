package service

import (
	"context"
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

func (s *sIndex) GetCarousels(ctx context.Context) ([]v1.IndexCarouselItem, error) {
	var carousels []v1.IndexCarouselItem
	err := dao.Carousel.Ctx(ctx).Where(do.Carousel{
		IsDeleted:  0,
	}).Scan(&carousels)
	return carousels, err
}

func (s *sIndex) GetHotGoods(ctx context.Context) ([]v1.GoodsItem, error) {
	var goods []v1.GoodsItem
	err := dao.IndexConfig.Ctx(ctx).Where(do.IndexConfig{
		ConfigType: 3,
		IsDeleted:  0,
	}).Scan(&goods)
	return goods, err
}

func (s *sIndex) GetNewGoods(ctx context.Context) ([]v1.GoodsItem, error) {
	var goods []v1.GoodsItem
	err := dao.IndexConfig.Ctx(ctx).Where(do.IndexConfig{
		ConfigType: 4,
		IsDeleted:  0,
	}).Scan(&goods)
	return goods, err
}

func (s *sIndex) GetRecommendGoods(ctx context.Context) ([]v1.GoodsItem, error) {
	var goods []v1.GoodsItem
	err := dao.IndexConfig.Ctx(ctx).Where(do.IndexConfig{
		ConfigType: 5,
		IsDeleted:  0,
	}).Scan(&goods)
	return goods, err
}
