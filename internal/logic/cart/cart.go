package cart

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"mall-api/internal/dao"
	"mall-api/internal/model"
	"mall-api/internal/service"
)

type (
	sCart struct{}
)

func init() {
	service.RegisterCart(New())
}

func New() *sCart {
	return &sCart{}
}

// GetUserCart 获取购物车列表
func (s *sCart) GetUserCart(ctx context.Context, userId int) (*[]model.UserCartOutput, error) {
	var cartRes = make([]model.UserCartOutput, 0)

	err :=
		g.Model("tb_newbee_mall_shopping_cart_item cart").
			InnerJoin("tb_newbee_mall_goods_info goods", "cart.goods_id=goods.goods_id").
			Fields("cart.cart_item_id,cart.goods_count,goods.goods_cover_img,goods.goods_name,goods.selling_price").
			Where(g.Map{"cart.is_deleted": 0, "cart.user_id": userId}).
			Scan(&cartRes)
	if err != nil {
		return nil, err
	}
	return &cartRes, err
}

// DelShopCart 删除购物车中商品
func (s *sCart) DelShopCart(ctx context.Context, userId int, cartId string) bool {
	shoppingCartItemCls := dao.ShoppingCartItem.Columns()
	_, err := dao.ShoppingCartItem.Ctx(ctx).
		Data(shoppingCartItemCls.IsDeleted, 1).
		Where(shoppingCartItemCls.UserId, userId).
		Where(shoppingCartItemCls.IsDeleted, 0).
		Where(shoppingCartItemCls.CartItemId, cartId).
		Update()
	if err != nil {
		return false
	}
	return true
}

// AddShopCart 新增商品到购物车
func (s *sCart) AddShopCart(ctx context.Context, userId int, goodsId int, goodsCount int) (*model.AddCartOutput, error) {
	// 判断商品是否存在
	_, err := service.Goods().Detail(ctx, gconv.String(goodsId))
	if err != nil {
		return nil, err
	}
	// 判断商品是否已在购物车中
	shoppingCartItemCls := dao.ShoppingCartItem.Columns()
	exists, err := dao.ShoppingCartItem.Ctx(ctx).
		Where(shoppingCartItemCls.IsDeleted, 0).
		Where(shoppingCartItemCls.GoodsId, goodsId).
		Where(shoppingCartItemCls.UserId, userId).
		One()
	if err != nil {
		return nil, err
	}
	if exists != nil {
		return nil, gerror.New("商品已在购物车中")
	}
	_, err = dao.ShoppingCartItem.Ctx(ctx).
		Data(g.Map{
			shoppingCartItemCls.UserId:     userId,
			shoppingCartItemCls.GoodsId:    goodsId,
			shoppingCartItemCls.GoodsCount: goodsCount,
		}).
		Insert()
	if err != nil {
		return nil, err
	}
	return &model.AddCartOutput{}, nil
}

// UpdateShopCart 修改购物车中商品数量
func (s *sCart) UpdateShopCart(ctx context.Context, userId int, cartItemId int, goodsCount int) (*model.UpdateCartOutput, error) {
	shoppingCartItemCls := dao.ShoppingCartItem.Columns()
	_, err := dao.ShoppingCartItem.Ctx(ctx).
		Data(shoppingCartItemCls.GoodsCount, goodsCount).
		Where(shoppingCartItemCls.IsDeleted, 0).
		Where(shoppingCartItemCls.CartItemId, cartItemId).
		Where(shoppingCartItemCls.UserId, userId).
		Update()
	if err != nil {
		return nil, err
	}
	return &model.UpdateCartOutput{}, nil
}

// GetCartGoodsById 根据购物车ID获取购物商品信息
func (s *sCart) GetCartGoodsById(ctx context.Context, cartItemId int) (*model.CartSettleOutput, error) {
	var cartRes model.CartSettleOutput
	err := g.Model("tb_newbee_mall_shopping_cart_item cart").
		InnerJoin("tb_newbee_mall_goods_info goods", "cart.goods_id=goods.goods_id").
		Fields("cart.cart_item_id,cart.goods_count,goods.goods_id,goods.goods_cover_img,goods.goods_name,goods.selling_price").
		Where(g.Map{"cart.is_deleted": 0, "cart.cart_item_id": cartItemId}).
		Scan(&cartRes)
	if err != nil {
		return nil, err
	}
	return &cartRes, err
}

// CartSettle 结算
func (s *sCart) CartSettle(ctx context.Context, cartItemIds []string) (res []model.CartSettleOutput, err error) {
	cartGoods := make([]model.CartSettleOutput, 0)
	for _, cartItemId := range cartItemIds {
		if cartItem, err := s.GetCartGoodsById(ctx, gconv.Int(cartItemId)); err == nil {
			cartGoods = append(cartGoods, *cartItem)
		}
	}
	return cartGoods, nil
}
