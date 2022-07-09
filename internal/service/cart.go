package service

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "mall-api/api/v1"
	"mall-api/internal/dao"
)

type (
	sCart struct{}
)

var (
	insCart = sCart{}
)

func Cart() *sCart {
	return &insCart
}

/**
获取购物车列表
*/
func (s *sCart) GetUserCart(ctx context.Context, userId int) (*[]v1.GetCartRes, error) {
	var cartRes []v1.GetCartRes

	err := g.Model("tb_newbee_mall_shopping_cart_item cart").InnerJoin("tb_newbee_mall_goods_info goods", "cart.goods_id=goods.goods_id").Fields(
		"cart.cart_item_id,cart.goods_count,goods.goods_cover_img,goods.goods_name,goods.selling_price",
	).Where(g.Map{"cart.is_deleted": 0, "cart.user_id": userId}).Scan(&cartRes)
	if err != nil {
		return nil, err
	}
	return &cartRes, err
}

/**
删除购物车中商品
*/
func (s *sCart) DelShopCart(ctx context.Context, userId int, cartId string) bool {
	_, err := g.Model("tb_newbee_mall_shopping_cart_item").Data(g.Map{"is_deleted": 1}).Where(g.Map{"is_deleted": 0, "cart_item_id": cartId, "user_id": userId}).Update()
	if err != nil {
		return false
	}
	return true
}

/**
新增商品到购物车
*/
func (s *sCart) AddShopCart(ctx context.Context, userId int, goodsId int, goodsCount int) (*v1.AddCartRes, error) {
	// 判断商品是否存在
	_, err := Goods().Detail(ctx, gconv.String(goodsId))
	if err != nil {
		return nil, err
	}
	// 判断商品是否已在购物车中
	exists, err := g.Model("tb_newbee_mall_shopping_cart_item").Where(g.Map{"is_deleted": 0, "goods_id": goodsId, "user_id": userId}).One()
	if err != nil {
		return nil, err
	}
	if exists!=nil{
		return nil, gerror.New("商品已在购物车中")
	}
	_, err = dao.ShoppingCartItem.Ctx(ctx).Insert(g.Map{"user_id": userId, "goods_id": goodsId, "goods_count": goodsCount})
	if err != nil {
		return nil, err
	}
	return &v1.AddCartRes{}, nil
}

/**
修改购物车中商品数量
 */
func (s *sCart) UpdateShopCart(ctx context.Context, userId int, cartItemId int, goodsCount int) (*v1.UpdateCartRes, error) {
	_, err := g.Model("tb_newbee_mall_shopping_cart_item").Data(g.Map{"goods_count": goodsCount}).Where(g.Map{"is_deleted": 0, "cart_item_id": cartItemId, "user_id": userId}).Update()
	if err != nil {
		return nil, err
	}
	return &v1.UpdateCartRes{}, nil
}

/**
根据购物车ID获取购物商品信息
 */
func (s *sCart) GetCartGoodsById(ctx context.Context, cartItemId int) (*v1.CartSettleRes, error) {
	var cartRes v1.CartSettleRes
	err := g.Model("tb_newbee_mall_shopping_cart_item cart").InnerJoin("tb_newbee_mall_goods_info goods", "cart.goods_id=goods.goods_id").Fields(
		"cart.cart_item_id,cart.goods_count,goods.goods_id,goods.goods_cover_img,goods.goods_name,goods.selling_price",
	).Where(g.Map{"cart.is_deleted": 0, "cart.cart_item_id": cartItemId}).Scan(&cartRes)
	if err != nil {
		return nil, err
	}
	return &cartRes, err
}