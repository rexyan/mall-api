package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "mall-api/api/v1"
	"mall-api/internal/model"
	"mall-api/internal/service"
)

var Search = cSearch{}

type cSearch struct {
}

func (c *cSearch) IndexSearch(ctx context.Context, req *v1.IndexSearchReq) (res *v1.IndexSearchRes, err error) {
	indexData, err := service.Search().IndexSearch(ctx, req.Keyword, req.GoodsCategoryId, req.OrderBy, model.PageReq{
		PageNumber: req.PageNumber,
		PageSize:   req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	if err := gconv.Structs(res, &indexData); err != nil {
		return nil, err
	}
	return
}
