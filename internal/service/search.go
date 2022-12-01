// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package service

import (
	"context"
	"mall-api/internal/model"
)

type ISearch interface {
	IndexSearch(ctx context.Context, keyword string, categoryId int, orderBy string, page model.PageReq) (*model.IndexSearchOutput, error)
}

var localSearch ISearch

func Search() ISearch {
	if localSearch == nil {
		panic("implement not found for interface ISearch, forgot register?")
	}
	return localSearch
}

func RegisterSearch(i ISearch) {
	localSearch = i
}
