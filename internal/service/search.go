package service

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	v1 "mall-api/api/v1"
	"mall-api/internal/dao"
)

type sSearch struct {
}

var insSearch = sSearch{}

func Search() *sSearch {
	return &insSearch
}

func (s *sSearch) IndexSearch(ctx context.Context, keyword string, categoryId int, orderBy string, page v1.PageReq) (*v1.IndexSearchRes, error) {
	r := g.RequestFromCtx(ctx)
	var IndexSearchGoodsInfo v1.IndexSearchRes
	IndexSearchGoodsInfo.List = make([]v1.SearchGoodsItem, 0)

	// 排序
	orderByCol := "create_time"
	if orderBy == "price"{
		orderByCol = "selling_price"
	}
	// 条件过滤
	condition := g.Map{
		"goods_name like '%?%'": keyword,
	}
	if categoryId >=0 {
		condition["goods_category_id"] = categoryId
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
