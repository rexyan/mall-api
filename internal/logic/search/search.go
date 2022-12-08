package search

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"mall-api/internal/dao"
	"mall-api/internal/model"
	"mall-api/internal/service"
)

type sSearch struct {
}

func init() {
	service.RegisterSearch(New())
}

func New() *sSearch {
	return &sSearch{}
}

func (s *sSearch) IndexSearch(ctx context.Context, keyword string, categoryId int, orderBy string, page model.PageReq) (*model.IndexSearchOutput, error) {
	r := g.RequestFromCtx(ctx)
	var IndexSearchGoodsInfo model.IndexSearchOutput
	IndexSearchGoodsInfo.List = make([]model.SearchGoodsItem, 0)

	goodsInfoCls := dao.GoodsInfo.Columns()
	// 排序
	orderByCol := goodsInfoCls.CreateTime
	if orderBy == "price" {
		orderByCol = "selling_price"
	}
	// 条件过滤
	condition := g.Map{
		"goods_name like '%?%'": keyword,
	}
	if categoryId >= 0 {
		condition[goodsInfoCls.GoodsCategoryId] = categoryId
	}

	query := dao.GoodsInfo.Ctx(ctx).Where(condition).OrderDesc(orderByCol)

	// 总数
	count, err := query.Count()
	if err != nil {
		return nil, err
	}
	// 分页信息
	pageInfo := r.GetPage(count, page.PageSize)
	// 分页数据
	err = query.Page(page.PageNumber, page.PageSize).Scan(&IndexSearchGoodsInfo.List)
	if err != nil {
		return nil, err
	}
	IndexSearchGoodsInfo.CurrentPage = pageInfo.CurrentPage
	IndexSearchGoodsInfo.PageBarNum = pageInfo.PageBarNum
	IndexSearchGoodsInfo.TotalSize = pageInfo.TotalSize
	IndexSearchGoodsInfo.TotalPage = pageInfo.TotalPage
	return &IndexSearchGoodsInfo, err
}
