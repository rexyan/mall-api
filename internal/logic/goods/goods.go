package goods

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"mall-api/internal/dao"
	"mall-api/internal/model"
	"mall-api/internal/model/entity"
	"mall-api/internal/service"
)

type (
	sGoods struct{}
)

func init() {
	service.RegisterGoods(New())
}

func New() *sGoods {
	return &sGoods{}
}

// Detail 获取商品详情
func (s *sGoods) Detail(ctx context.Context, goodId string) (*model.GoodsDetailOutput, error) {
	var goodsInfo *entity.GoodsInfo
	err := dao.GoodsInfo.Ctx(ctx).Where(g.Map{"goods_id": goodId}).Scan(&goodsInfo)
	if err != nil {
		return nil, err
	}
	if goodsInfo == nil {
		return nil, gerror.New("未找到商品信息!")
	}
	return &model.GoodsDetailOutput{
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
