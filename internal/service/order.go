package service

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/guuid"
	v1 "mall-api/api/v1"
)

type sOrder struct {
}

var insOrder = sOrder{}

func Order() *sOrder {
	return &insOrder
}

/**
获取下单时购物车商品信息
 */
func (s *sOrder) GetOrderGoodsInfos(ctx context.Context, cartItemIds []int) (*[]v1.CartSettleRes, error) {
	// 查询商品信息
	var orderGoodsInfos []v1.CartSettleRes
	for _, cartItemId := range cartItemIds {
		goodsInfo, err := Cart().GetCartGoodsById(ctx, cartItemId)
		if err == nil {
			orderGoodsInfos = append(orderGoodsInfos, *goodsInfo)
		}
	}
	return &orderGoodsInfos, nil
}

/**
下单
 */
func (s *sOrder) SaveOrder(ctx context.Context, userId string, addressId int, cartItemIds []int) error {
	// 生成订单 ID
	uuid, err := guuid.NewRandom()
	if err != nil {
		return gerror.New("下单失败!")
	}
	orderNo := gstr.SubStr(gconv.String(uuid), 0, 20)

	err = g.DB("default").Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 查询地址信息
		address, err := Address().GetAddressById(ctx, gconv.String(addressId))
		if err != nil {
			return err
		}

		orderGoodsInfos, err := Order().GetOrderGoodsInfos(ctx, cartItemIds)
		if err != nil {
			return err
		}

		// 计算订单总价
		var totalPrice = 0
		for _, goodsInfo := range *orderGoodsInfos {
			totalPrice += goodsInfo.SellingPrice * goodsInfo.GoodsCount
		}

		// 生成订单
		order, err := tx.Ctx(ctx).Insert("tb_newbee_mall_order", g.Map{
			"order_no":     orderNo,
			"user_id":      userId,
			"total_price":  totalPrice,
			"pay_status":   0, // 支付状态:0.未支付,1.支付成功,-1:支付失败
			"pay_type":     0, // 0.无 1.支付宝支付 2.微信支付
			"order_status": 0, // 订单状态:0.待支付 1.已支付 2.配货完成 3:出库成功 4.交易成功 -1.手动关闭 -2.超时关闭 -3.商家关闭
			"extra_info":   "",
		})
		if err != nil {
			return err
		}

		// 获取订单 ID
		orderId, err := order.LastInsertId()
		if err != nil {
			return err
		}

		// 订单关联地址
		_, err = tx.Ctx(ctx).Insert("tb_newbee_mall_order_address", g.Map{
			"order_id":       orderId,
			"user_name":      address.UserName,
			"user_phone":     address.UserPhone,
			"province_name":  address.ProvinceName,
			"city_name":      address.CityName,
			"region_name":    address.RegionName,
			"detail_address": address.DetailAddress,
		})
		if err != nil {
			return err
		}

		// 订单商品快照
		for _, goodsInfo := range *orderGoodsInfos {
			_, err = tx.Ctx(ctx).Insert("tb_newbee_mall_order_item", g.Map{
				"order_id":        orderId,
				"goods_id":        goodsInfo.GoodsId,
				"goods_name":      goodsInfo.GoodsName,
				"goods_cover_img": goodsInfo.GoodsCoverImg,
				"selling_price":   goodsInfo.SellingPrice,
				"goods_count":     goodsInfo.GoodsCount,
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	return nil
}
