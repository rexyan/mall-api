package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"mall-api/internal/service"

	"mall-api/api/v1"
)

var (
	Index = cIndex{}
)

type cIndex struct{}

func (c *cIndex) IndexInfos(ctx context.Context, req *v1.IndexInfosReq) (res *v1.IndexInfosRes, err error) {
	var IndexCarouselItem = &[]v1.IndexCarouselItem{}
	var IndexHotGoods = &[]v1.GoodsItem{}
	var IndexNewGoods = &[]v1.GoodsItem{}
	var IndexRecommendGoods = &[]v1.GoodsItem{}

	carousels, err := service.Index().GetCarousels(ctx)
	if err != nil {
		return nil, err
	}
	if err := gconv.Scan(IndexCarouselItem, &carousels); err != nil {
		return nil, err
	}

	hotGoods, err := service.Index().GetHotGoods(ctx)
	if err != nil {
		return nil, err
	}
	if err := gconv.Scan(IndexHotGoods, &hotGoods); err != nil {
		return nil, err
	}

	newGoods, err := service.Index().GetNewGoods(ctx)
	if err != nil {
		return nil, err
	}
	if err := gconv.Scan(IndexNewGoods, &newGoods); err != nil {
		return nil, err
	}

	recommendGoods, err := service.Index().GetRecommendGoods(ctx)
	if err != nil {
		return nil, err
	}
	if err := gconv.Scan(IndexRecommendGoods, &recommendGoods); err != nil {
		return nil, err
	}

	return &v1.IndexInfosRes{
		IndexCarouselItem:   *IndexCarouselItem,
		IndexHotGoods:       *IndexHotGoods,
		IndexNewGoods:       *IndexNewGoods,
		IndexRecommendGoods: *IndexRecommendGoods,
	}, nil
}
