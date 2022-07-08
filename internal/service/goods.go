package service

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "mall-api/api/v1"
	"mall-api/internal/dao"
	"mall-api/internal/model/entity"
)

type (
	sGoods struct{}
)

var (
	insGoods = sGoods{}
)

func Goods() *sGoods {
	return &insGoods
}

/**
获取商品详情
*/
func (s *sGoods) Detail(ctx context.Context, goodId string) (*v1.GoodsDetailRes, error) {
	var goodsInfo *entity.GoodsInfo
	err := dao.GoodsInfo.Ctx(ctx).Where(g.Map{"goods_id": goodId}).Scan(&goodsInfo)
	if err != nil {
		return nil, err
	}
	if goodsInfo == nil{
		return nil, gerror.New("未找到商品信息!")
	}
	return &v1.GoodsDetailRes{
		GoodsCarouselList:  gconv.Strings(goodsInfo.GoodsCarousel),
		GoodsCoverImg:      goodsInfo.GoodsCoverImg,
		GoodsDetailContent: goodsInfo.GoodsDetailContent,
		GoodsId:            gconv.Int(goodsInfo.GoodsId),
		GoodsIntro:         goodsInfo.GoodsIntro,
		GoodsName:          goodsInfo.GoodsName,
		OriginalPrice:      goodsInfo.OriginalPrice,
		SellingPrice:       goodsInfo.SellingPrice,
		Tag:                goodsInfo.Tag,
	}, nil
}
