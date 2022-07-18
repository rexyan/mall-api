package controller

import (
	"context"
	v1 "mall-api/api/v1"
	"mall-api/internal/service"
)

var Search = cSearch{}

type cSearch struct {

}

func (c *cSearch) IndexSearch(ctx context.Context, req *v1.IndexSearchReq) (*v1.IndexSearchRes, error)  {
	return service.Search().IndexSearch(ctx, req.Keyword, req.GoodsCategoryId, req.OrderBy, req.PageReq)
}
