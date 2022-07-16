package service

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "mall-api/api/v1"
	"mall-api/internal/dao"
	"mall-api/internal/model/entity"
)

type (
	sCategory struct{}
)

var (
	insCategory = sCategory{}
)

func Category() *sCategory {
	return &insCategory
}

/**
根据 ParentId 获取 Category 信息
*/
func (s *sCategory) GetCategoryByParentId(ctx context.Context, parentId int) (*[]entity.GoodsCategory, error) {
	var goodsCategory []entity.GoodsCategory
	err := dao.GoodsCategory.Ctx(ctx).Where("parent_id", parentId).Scan(&goodsCategory)
	if err != nil {
		return nil, err
	}
	return &goodsCategory, nil
}

/**
根据 ID 获取 Category 信息
*/
func (s *sCategory) GetCategoryById(ctx context.Context, categoryId int) (*entity.GoodsCategory, error) {
	var goodsCategory entity.GoodsCategory
	err := dao.GoodsCategory.Ctx(ctx).Where("category_id", categoryId).Scan(&goodsCategory)
	if err != nil {
		return nil, err
	}
	return &goodsCategory, nil
}

/**
根据 Level 获取 Category 信息
*/
func (s *sCategory) GetCategoryByLevel(ctx context.Context, categoryLevel int) (*[]entity.GoodsCategory, error) {
	var goodsCategory []entity.GoodsCategory
	err := dao.GoodsCategory.Ctx(ctx).Where("category_level", categoryLevel).Scan(&goodsCategory)
	if err != nil {
		return nil, err
	}
	return &goodsCategory, nil
}

/**
获取 Category 信息
*/
func (s *sCategory) GetCategory(ctx context.Context) (*[]v1.CategoryRes, error) {
	var res []interface{}
	var categoryArray []v1.CategoryRes

	Category().GetCategoryTree(ctx, &res, 0)
	_ = gconv.Structs(res, &categoryArray)
	return &categoryArray, nil
}

func (s *sCategory) GetCategoryTree(ctx context.Context, res *[]interface{}, parentId int) {
	// 当前所有 parentId 的 Category
	category, _ := Category().GetCategoryByParentId(ctx, parentId)
	for _, c := range *category {

		var child = &g.Slice{}
		var item map[string]interface{}
		_ = gconv.Scan(c, &item)

		if c.CategoryLevel == 1 {
			item["secondLevelCategoryVOS"] = child
		} else if c.CategoryLevel == 2 {
			item["thirdLevelCategoryVOS"] = child
		}

		*res = append(*res, &item)
		Category().GetCategoryTree(ctx, child, gconv.Int(c.CategoryId))
	}
}
