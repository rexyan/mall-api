package v1

import "github.com/gogf/gf/v2/frame/g"

type CategoryItem struct {
	ParentId      string `json:"parentId"`
	CategoryId    string `json:"categoryId"`
	CategoryLevel string `json:"categoryLevel"`
	CategoryName  string `json:"categoryName"`
}

type CategoryReq struct {
	g.Meta        `path:"categories" tags:"分类" method:"get" summary:"分类"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
}

type CRes []CategoryRes

type CategoryRes struct {
	CategoryItem
	SecondLevelCategoryVOS []SecondLevelCategoryVOS `json:"secondLevelCategoryVOS"`
}

type SecondLevelCategoryVOS struct {
	CategoryItem
	ThirdLevelCategoryVOS []ThirdLevelCategoryVOS `json:"thirdLevelCategoryVOS"`
}

type ThirdLevelCategoryVOS struct {
	CategoryItem
}
