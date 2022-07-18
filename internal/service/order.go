package service

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/guuid"
	v1 "mall-api/api/v1"
	"mall-api/internal/dao"
	"mall-api/internal/model/entity"
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
func (s *sOrder) SaveOrder(ctx context.Context, userId string, addressId int, cartItemIds []int) (string, error) {
	// 生成订单 ID
	uuid, err := guuid.NewRandom()
	if err != nil {
		return "", gerror.New("下单失败!")
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
		// 删除购物车中本次下单的商品
		_, err = tx.Ctx(ctx).Update("tb_newbee_mall_shopping_cart_item", g.Map{
			"is_deleted": 1,
		}, g.Map{
			"cart_item_id in (?)": cartItemIds,
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	return orderNo, nil
}

/**
查询订单中的商品快照信息
*/
func (s *sOrder) GetOrderGoodsSnapshotById(ctx context.Context, orderId int) ([]v1.GetCartRes, error) {
	var snapshot []v1.GetCartRes
	err := dao.OrderItem.Ctx(ctx).Where("order_id", orderId).Scan(&snapshot)
	if err != nil {
		return nil, err
	}
	return snapshot, nil
}

/**
查询订单中的商品快照信息
*/
func (s *sOrder) GetOrderIdByNo(ctx context.Context, orderNo string) (orderId int, err error) {
	var order *entity.Order
	err = dao.Order.Ctx(ctx).Where("order_no", orderNo).Scan(&order)
	if err != nil {
		return 0, err
	}
	if order != nil {
		return gconv.Int(order.OrderId), nil
	}
	return 0, gerror.New("order not found!")
}

/**
查询用户订单
*/
func (s *sOrder) GetOrderByUser(ctx context.Context, userId string, orderStatus int, page v1.PageReq) (*v1.OrderListRes, error) {
	var orderList v1.OrderListRes
	orderList.List = make([]v1.OrderListItem, 0)

	r := g.RequestFromCtx(ctx)
	condition := map[string]interface{}{
		"user_id": userId,
	}
	if orderStatus >= 0 {
		condition["order_status"] = orderStatus
	}
	query := dao.Order.Ctx(ctx).Where(condition)
	// 总数
	count, err := query.Count()
	if err != nil {
		return nil, err
	}
	// 分页信息
	pageInfo := r.GetPage(count, page.PageSize)
	// 分页数据
	err = query.Page(page.PageNumber, page.PageSize).Scan(&orderList.List)
	if err != nil {
		return nil, err
	}
	// 获取订单中每个商品的信息

	for index, order := range orderList.List {
		snapshot, err := Order().GetOrderGoodsSnapshotById(ctx, order.OrderId)
		if err == nil {
			for _, s := range snapshot {
				orderList.List[index].NewBeeMallOrderItemVOS = append(orderList.List[index].NewBeeMallOrderItemVOS, s)
			}
		}
	}

	orderList.CurrentPage = pageInfo.CurrentPage
	orderList.PageBarNum = pageInfo.PageBarNum
	orderList.TotalSize = pageInfo.TotalSize
	orderList.TotalPage = pageInfo.TotalPage
	return &orderList, err
}

/**
查询订单详情
*/
func (s *sOrder) GetOrderDetail(ctx context.Context, orderNo string) (*v1.OrderDetailRes, error) {
	var orderDetail v1.OrderDetailRes
	// 根据订单编号查询订单 ID
	orderId, err := Order().GetOrderIdByNo(ctx, orderNo)
	if err != nil {
		return nil, err
	}
	// 查询订单商品快照信息
	snapshot, err := Order().GetOrderGoodsSnapshotById(ctx, orderId)
	err = dao.Order.Ctx(ctx).Where(g.Map{
		"order_no": orderNo,
	}).Scan(&orderDetail)
	if err == nil {
		for _, s := range snapshot {
			orderDetail.NewBeeMallOrderItemVOS = append(orderDetail.NewBeeMallOrderItemVOS, s)
		}
	}
	return &orderDetail, err
}

/**
订单支付
*/
func (s *sOrder) PayOrder(ctx context.Context, orderNo string, payType string) (*v1.PayOrderRes, error) {
	order, err := Order().GetOrderDetail(ctx, orderNo)
	if err != nil || order == nil {
		return nil, gerror.New("order not found!")
	}
	if order.OrderStatus != 0 {
		return nil, gerror.New("order status error!")
	}
	_, err = dao.Order.Ctx(ctx).Data(g.Map{
		"pay_type":     gconv.Int(payType),
		"pay_status":   1, // 支付成功
		"pay_time":     gtime.Datetime(),
		"order_status": 1, // 已支付
	}).Where("order_no", orderNo).Update()
	if err != nil {
		return nil, err
	}
	return &v1.PayOrderRes{}, err
}
