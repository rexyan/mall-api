package controller

import (
	"context"
	"mall-api/internal/service"

	"mall-api/api/v1"
)

var (
	Index = cIndex{}
)

type cIndex struct{}

func (c *cIndex) IndexInfos(ctx context.Context, req *v1.IndexInfosReq) (res *v1.IndexInfosRes, err error) {
	carousels, err := service.Index().GetCarousels(ctx)
	if err != nil {
		return nil, err
	}
	hotGoods, err := service.Index().GetHotGoods(ctx)
	if err != nil {
		return nil, err
	}
	newGoods, err := service.Index().GetNewGoods(ctx)
	if err != nil {
		return nil, err
	}
	recommendGoods, err := service.Index().GetRecommendGoods(ctx)
	if err != nil {
		return nil, err
	}

	return &v1.IndexInfosRes{
		IndexCarouselItem:   carousels,
		IndexHotGoods:       hotGoods,
		IndexNewGoods:       newGoods,
		IndexRecommendGoods: recommendGoods,
	}, nil
}
