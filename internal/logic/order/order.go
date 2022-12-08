package order

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/guuid"
	"mall-api/internal/dao"
	"mall-api/internal/model"
	"mall-api/internal/model/entity"
	"mall-api/internal/service"
)

type sOrder struct {
}

func init() {
	service.RegisterOrder(New())
}

func New() *sOrder {
	return &sOrder{}
}

// GetOrderGoodsInfos 获取下单时购物车商品信息
func (s *sOrder) GetOrderGoodsInfos(ctx context.Context, cartItemIds []int) (*[]model.CartSettleOutput, error) {
	// 查询商品信息
	var orderGoodsInfos []model.CartSettleOutput
	for _, cartItemId := range cartItemIds {
		goodsInfo, err := service.Cart().GetCartGoodsById(ctx, cartItemId)
		if err == nil {
			orderGoodsInfos = append(orderGoodsInfos, *goodsInfo)
		}
	}
	return &orderGoodsInfos, nil
}

// SaveOrder 下单
func (s *sOrder) SaveOrder(ctx context.Context, userId string, addressId int, cartItemIds []int) (string, error) {
	// 生成订单 ID
	uuid, err := guuid.NewRandom()
	if err != nil {
		return "", gerror.New("下单失败!")
	}
	orderNo := gstr.SubStr(gconv.String(uuid), 0, 20)

	err = g.DB("default").Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 查询地址信息
		address, err := service.Address().GetAddressById(ctx, gconv.String(addressId))
		if err != nil {
			return err
		}

		orderGoodsInfos, err := service.Order().GetOrderGoodsInfos(ctx, cartItemIds)
		if err != nil {
			return err
		}

		// 计算订单总价
		var totalPrice = 0
		for _, goodsInfo := range *orderGoodsInfos {
			totalPrice += goodsInfo.SellingPrice * goodsInfo.GoodsCount
		}

		// 订单字段
		orderCls := dao.Order.Columns()

		// 生成订单
		order, err := tx.Ctx(ctx).Insert("tb_newbee_mall_order", g.Map{
			orderCls.OrderNo:     orderNo,
			orderCls.UserId:      userId,
			orderCls.TotalPrice:  totalPrice,
			orderCls.PayStatus:   0, // 支付状态:0.未支付,1.支付成功,-1:支付失败
			orderCls.PayType:     0, // 0.无 1.支付宝支付 2.微信支付
			orderCls.OrderStatus: 0, // 订单状态:0.待支付 1.已支付 2.配货完成 3:出库成功 4.交易成功 -1.手动关闭 -2.超时关闭 -3.商家关闭
			orderCls.ExtraInfo:   "",
		})
		if err != nil {
			return err
		}

		// 获取订单 ID
		orderId, err := order.LastInsertId()
		if err != nil {
			return err
		}

		// 订单地址字段
		orderAddressCls := dao.OrderAddress.Columns()

		// 订单关联地址
		_, err = tx.Ctx(ctx).Insert("tb_newbee_mall_order_address", g.Map{
			orderAddressCls.OrderId:       orderId,
			orderAddressCls.UserName:      address.UserName,
			orderAddressCls.UserPhone:     address.UserPhone,
			orderAddressCls.ProvinceName:  address.ProvinceName,
			orderAddressCls.CityName:      address.CityName,
			orderAddressCls.RegionName:    address.RegionName,
			orderAddressCls.DetailAddress: address.DetailAddress,
		})
		if err != nil {
			return err
		}

		// 订单地址字段
		orderItemCls := dao.OrderItem.Columns()

		// 订单商品快照
		for _, goodsInfo := range *orderGoodsInfos {
			_, err = tx.Ctx(ctx).Insert("tb_newbee_mall_order_item", g.Map{
				orderItemCls.OrderId:       orderId,
				orderItemCls.GoodsId:       goodsInfo.GoodsId,
				orderItemCls.GoodsName:     goodsInfo.GoodsName,
				orderItemCls.GoodsCoverImg: goodsInfo.GoodsCoverImg,
				orderItemCls.SellingPrice:  goodsInfo.SellingPrice,
				orderItemCls.GoodsCount:    goodsInfo.GoodsCount,
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

// GetOrderGoodsSnapshotById 查询订单中的商品快照信息
func (s *sOrder) GetOrderGoodsSnapshotById(ctx context.Context, orderId int) ([]model.GetCartOutput, error) {
	var snapshot []model.GetCartOutput
	err := dao.OrderItem.Ctx(ctx).Where(dao.OrderItem.Columns().OrderId, orderId).Scan(&snapshot)
	if err != nil {
		return nil, err
	}
	return snapshot, nil
}

// GetOrderIdByNo 查询订单中的商品快照信息
func (s *sOrder) GetOrderIdByNo(ctx context.Context, orderNo string) (orderId int, err error) {
	var order *entity.Order
	err = dao.Order.Ctx(ctx).Where(dao.Order.Columns().OrderNo, orderNo).Scan(&order)
	if err != nil {
		return 0, err
	}
	if order != nil {
		return gconv.Int(order.OrderId), nil
	}
	return 0, gerror.New("order not found!")
}

// GetOrderByUser 查询用户订单
func (s *sOrder) GetOrderByUser(ctx context.Context, userId string, orderStatus int, page model.PageReq) (*model.OrderListOutput, error) {
	var orderList model.OrderListOutput
	orderList.List = make([]model.OrderListItem, 0)

	r := g.RequestFromCtx(ctx)
	condition := map[string]interface{}{
		dao.User.Columns().UserId: userId,
	}
	if orderStatus >= 0 {
		condition[dao.Order.Columns().OrderStatus] = orderStatus
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
		snapshot, err := service.Order().GetOrderGoodsSnapshotById(ctx, order.OrderId)
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

// GetOrderDetail 查询订单详情
func (s *sOrder) GetOrderDetail(ctx context.Context, orderNo string) (*model.OrderDetailOutput, error) {
	var orderDetail model.OrderDetailOutput
	// 根据订单编号查询订单 ID
	orderId, err := service.Order().GetOrderIdByNo(ctx, orderNo)
	if err != nil {
		return nil, err
	}
	// 查询订单商品快照信息
	snapshot, err := service.Order().GetOrderGoodsSnapshotById(ctx, orderId)
	err = dao.Order.Ctx(ctx).
		Where(dao.Order.Columns().OrderNo, orderNo).
		Scan(&orderDetail)
	if err == nil {
		for _, s := range snapshot {
			orderDetail.NewBeeMallOrderItemVOS = append(orderDetail.NewBeeMallOrderItemVOS, s)
		}
	}
	return &orderDetail, err
}

// PayOrder 订单支付
func (s *sOrder) PayOrder(ctx context.Context, orderNo string, payType string) (*model.PayOrderOutput, error) {
	order, err := service.Order().GetOrderDetail(ctx, orderNo)
	if err != nil || order == nil {
		return nil, gerror.New("order not found!")
	}
	if order.OrderStatus != 0 {
		return nil, gerror.New("order status error!")
	}
	orderCls := dao.Order.Columns()
	_, err = dao.Order.Ctx(ctx).Data(g.Map{
		orderCls.PayType:     gconv.Int(payType),
		orderCls.PayStatus:   1, // 支付成功
		orderCls.PayTime:     gtime.Datetime(),
		orderCls.OrderStatus: 1, // 已支付
	}).Where(orderCls.OrderNo, orderNo).Update()
	if err != nil {
		return nil, err
	}
	return &model.PayOrderOutput{}, err
}
