package service

import (
	"context"
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
func (s *sCategory) GetCategoryByParentId(ctx context.Context, parentId int) (*entity.GoodsCategory, error) {
	var goodsCategory *entity.GoodsCategory
	err := dao.GoodsCategory.Ctx(ctx).Where("parent_id", parentId).Scan(goodsCategory)
	if err!=nil{
		return nil, err
	}
	return goodsCategory, nil
}

/**
根据 ID 获取 Category 信息
 */
func (s *sCategory) GetCategoryById(ctx context.Context, categoryId int) (*[]entity.GoodsCategory, error) {
	var goodsCategory *[]entity.GoodsCategory
	err := dao.GoodsCategory.Ctx(ctx).Where("category_id", categoryId).Scan(goodsCategory)
	if err!=nil{
		return nil, err
	}
	return goodsCategory, nil
}

/**
根据 Level 获取 Category 信息
*/
func (s *sCategory) GetCategoryByLevel(ctx context.Context, categoryLevel int) (*[]entity.GoodsCategory, error) {
	var goodsCategory *[]entity.GoodsCategory
	err := dao.GoodsCategory.Ctx(ctx).Where("category_level", categoryLevel).Scan(goodsCategory)
	if err!=nil{
		return nil, err
	}
	return goodsCategory, nil
}

/**
获取 Category 信息
 */
func (s *sCategory) GetCategory(ctx context.Context) (*v1.CategoryRes, error) {
	var categoryRes  = v1.CategoryRes{}
	thirdLevelCategory, _ := Category().GetCategoryByLevel(ctx, 3)
	// secondLevelCategory, _ := Category().GetCategoryByLevel(ctx, 2)
	// firstLevelCategory, _ := Category().GetCategoryByLevel(ctx, 1)
	for _, thirdCategory := range *thirdLevelCategory {
		var secondLevelCategoryVOS  = v1.SecondLevelCategoryVOS{}
		var thirdLevelCategoryVOS  = v1.ThirdLevelCategoryVOS{}

		var thirdCategoryItem  *v1.CategoryItem
		// 获取二级分类信息
		secondLevelCategory, _ := Category().GetCategoryByParentId(ctx, gconv.Int(thirdCategory.CategoryId))
		// 获取二级本身的信息
		secondLevelCategoryInfo, _ := Category().GetCategoryById(ctx, gconv.Int(secondLevelCategory.CategoryId))
		_ = gconv.Scan(secondLevelCategoryVOS.CategoryItem, secondLevelCategoryInfo)

		// 获取一级分类信息
		firstLevelCategory, _ := Category().GetCategoryByParentId(ctx, gconv.Int(secondLevelCategory.ParentId))
		_ = gconv.Scan(categoryRes, firstLevelCategory)


		_ = gconv.Scan(thirdCategory, &thirdCategoryItem)
		thirdLevelCategoryVOS.CategoryItem = append(thirdLevelCategoryVOS.CategoryItem, *thirdCategoryItem)
	}
}
